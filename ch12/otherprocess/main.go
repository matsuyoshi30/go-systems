package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s [pid]\n", os.Args[0])
		return
	}

	// ほかプロセスID
	pid, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	// プロセスIDからos.Process取得
	process, err := os.FindProcess(pid)
	if err != nil {
		panic(err)
	}

	// シグナルを送る
	process.Signal(os.Kill)
	process.Kill() // killの場合はこんな感じでもかける
}
