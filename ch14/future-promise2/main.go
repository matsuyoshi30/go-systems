// 複数回値が参照できるFuture

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// チャネルをラップして、初回に取得した値をキャッシュし、2回目はキャッシュを返す
// -> 複数のタスクがFutureを参照できる
type StringFuture struct {
	receiver chan string
	cache    string
}

func NewStringFuture() (*StringFuture, chan string) {
	f := &StringFuture{
		receiver: make(chan string),
	}
	return f, f.receiver
}

func (f *StringFuture) Get() string {
	r, ok := <-f.receiver
	if ok {
		close(f.receiver)
		f.cache = r
	}
	return f.cache
}

func (f *StringFuture) Close() {
	close(f.receiver)
}

func readFile(path string) *StringFuture {
	promise, future := NewStringFuture() // ファイルを読み込んだ結果を返すFutureを返す
	go func() {
		content, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Printf("read error %s\n", err.Error())
			promise.Close()
		} else {
			future <- string(content)
		}
	}()
	return promise
}

func printFunc(futureSource *StringFuture) chan []string {
	promise := make(chan []string) // 文字列中の関数一覧を返すFutureを返す
	go func() {
		var result []string
		for _, line := range strings.Split(futureSource.Get(), "\n") { // futureが解決するまでループ
			if strings.HasPrefix(line, "func ") {
				result = append(result, line)
			}
		}
		promise <- result
	}()

	return promise
}

func countLines(futureSource *StringFuture) chan int {
	promise := make(chan int)
	go func() {
		promise <- len(strings.Split(futureSource.Get(), "\n"))
	}()
	return promise
}

func main() {
	futureSource := readFile("main.go")
	futureFuncs := printFunc(futureSource)
	fmt.Println(strings.Join(<-futureFuncs, "\n"))
	fmt.Println(<-countLines(futureSource))
}