package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2) // ジョブ数を登録

	go func() {
		fmt.Println("work 1")
		wg.Done() // 完了通知
	}()

	go func() {
		time.Sleep(time.Second * 3)
		fmt.Println("work 2")
		wg.Done() // 完了通知
	}()

	wg.Wait() // 登録された全てのジョブが完了するのを待つ
	fmt.Println("done")
}
