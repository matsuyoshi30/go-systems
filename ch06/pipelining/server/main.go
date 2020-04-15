package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

// conn に書き出し (順番に従って実行される)
func writeToConn(sessionResponses chan chan *http.Response, conn net.Conn) {
	defer conn.Close()

	for sessionResponse := range sessionResponses {
		response := <-sessionResponse
		response.Write(conn)
		close(sessionResponse)
	}
}

// セッション内の request を処理
func handleRequest(request *http.Request, resultReceiver chan *http.Response) {
	dump, err := httputil.DumpRequest(request, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))

	content := "Hello world\n"
	response := &http.Response{
		StatusCode:    200,
		ProtoMajor:    1,
		ProtoMinor:    1, // HTTP1.1
		ContentLength: int64(len(content)),
		Body:          ioutil.NopCloser(strings.NewReader(content)),
	}

	resultReceiver <- response // ブロックされていた writeToConn の処理が再起動される
}

// 1つのセッションを処理
func processSession(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())

	// request を順番に処理するためのチャネル
	sessionResponses := make(chan chan *http.Response, 50)
	defer close(sessionResponses)

	go writeToConn(sessionResponses, conn)
	reader := bufio.NewReader(conn)

	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))

		request, err := http.ReadRequest(reader)
		if err != nil {
			neterr, ok := err.(net.Error)
			if ok && neterr.Timeout() {
				fmt.Println("Timeout")
				break
			} else if err == io.EOF {
				break
			}
			panic(err)
		}

		sessionResponse := make(chan *http.Response)
		sessionResponses <- sessionResponse // バッファ付きチャネル内にバッファなしチャネルが！

		go handleRequest(request, sessionResponse)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running at localhost:8888")

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go processSession(conn)
	}
}
