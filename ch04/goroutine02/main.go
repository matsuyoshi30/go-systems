package main

import (
	"fmt"
	"time"
)

func sub() {
}

func main() {
	fmt.Println("start sub()")
	go func() {
		fmt.Println("sub() is running")
		time.Sleep(time.Second)
		fmt.Println("sub() is finished")
	}()
	time.Sleep(2 * time.Second) // これがないと sub() の完了前に main() が終わっちゃう
}
