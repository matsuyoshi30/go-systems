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

		go func() {
			defer conn.Close()

			fmt.Printf("Accept %v\n", conn.RemoteAddr())

			for {
				// set timeout
				conn.SetReadDeadline(time.Now().Add(5 * time.Second))

				request, err := http.ReadRequest(bufio.NewReader(conn))
				if err != nil {
					neterr, ok := err.(net.Error) // 型アサーション！ down cast
					if ok && neterr.Timeout() {
						fmt.Println("Timeout")
						break
					} else if err == io.EOF {
						break
					}
					panic(err)
				}

				dump, err := httputil.DumpRequest(request, true)
				if err != nil {
					panic(err)
				}
				fmt.Println(string(dump))

				content := "Hello world\n"
				response := http.Response{
					StatusCode:    200,
					ProtoMajor:    1,
					ProtoMinor:    1, // HTTP1.1
					ContentLength: int64(len(content)),
					Body:          ioutil.NopCloser(strings.NewReader(content)),
				}
				response.Write(conn)
			}
		}()
	}
}
