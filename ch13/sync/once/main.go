package main

import (
	"fmt"
	"sync"
)

func initialize() {
	fmt.Println("initialize process")
}

var once sync.Once

func main() {
	// 3回呼んでも1回しか実行されない
	once.Do(initialize)
	once.Do(initialize)
	once.Do(initialize)
}
