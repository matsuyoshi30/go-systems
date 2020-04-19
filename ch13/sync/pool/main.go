package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var count int
	pool := sync.Pool{
		New: func() interface{} { // 新規作成時
			count++
			return fmt.Sprintf("created: %d", count)
		},
	}

	pool.Put("manually added: 1")
	pool.Put("manually added: 2")

	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	fmt.Println(pool.Get()) // New

	pool.Put("removed 1")
	runtime.GC() // 保持していた removed 1 が削除される
	fmt.Println(pool.Get())
}
