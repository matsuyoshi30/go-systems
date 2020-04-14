package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	che := make(chan struct{})

	fmt.Println("start")

	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				ch1 <- i
			} else {
				ch2 <- i
			}
		}
		close(che)
	}()

	for {
		select {
		case c1 := <-ch1:
			fmt.Println("received from ch1", c1)
		case c2 := <-ch2:
			fmt.Println("received from ch2", c2)
		case <-che:
			fmt.Println("end")
			return
		default:
			break
		}
	}
}
