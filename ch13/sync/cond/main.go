package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex

	cond := sync.NewCond(&mutex)

	for _, name := range []string{"A", "B", "C"} {
		go func(name string) {
			// ロックしてからwait
			mutex.Lock()
			defer mutex.Unlock()
			// Broadcast()が呼ばれるまで待つ
			cond.Wait()
			fmt.Println(name)
		}(name)
	}

	fmt.Println("ready...")
	time.Sleep(time.Second)
	fmt.Println("Go!")
	// 待っているgoroutineを一斉に起こす
	cond.Broadcast()

	time.Sleep(time.Second) // これないとmainが先に終わっちゃうので説明用
}
