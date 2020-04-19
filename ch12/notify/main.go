package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signals := make(chan os.Signal, 1) // サイズが1より大きいチャネル

	// 最初は作成したチャネル、後は可変長引数で任意のシグナルを登録
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// シグナルはチャネルを通して受け取る
	s := <-signals
	switch s {
	case syscall.SIGINT:
		fmt.Println("SIGINT")
	case syscall.SIGTERM:
		fmt.Println("SIGTERM")
	}
}
