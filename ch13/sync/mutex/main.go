package main

import (
	"fmt"
	"sync"
)

var id int

func generateId(mutex *sync.Mutex) int {
	// Lock()/Unlock()はペアで呼び出して使う
	mutex.Lock()
	defer mutex.Unlock() // ペアで使用するものは連続して書くのがベストプラクティス（deferつかおう）
	id++
	return id
}

func main() {
	// ゼロ値で使用できる
	var mutex sync.Mutex

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Printf("id: %d\n", generateId(&mutex)) // idをインクリメントするコードが1つしか実行されない
		}()
	}
}
