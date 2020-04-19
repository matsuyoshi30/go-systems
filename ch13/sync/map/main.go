package main

import (
	"fmt"
	"sync"
)

func main() {
	smap := &sync.Map{}

	smap.Store("hello", "world")
	smap.Store(1, 2)
	smap.Store("test", "will delete")

	smap.Delete("test")

	value, ok := smap.Load("hello")
	fmt.Printf("key=%v value=%v exists?=%v\n", "hello", value, ok)

	smap.Range(func(key, value interface{}) bool {
		fmt.Printf("%v: %v\n", key, value)
		return true
	})

	smap.LoadOrStore(1, 3) // 既に登録されているキーなので無視
	smap.LoadOrStore(2, 4) // これは登録される
}
