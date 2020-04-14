package main

import (
	"fmt"
	"math"
)

func primeNumber() chan int {
	result := make(chan int)

	go func() {
		result <- 2
		for i := 3; i < 100000; i += 2 {
			// 整数iの平方根まで、自身を割り切る整数があるかを探索
			l := int(math.Sqrt(float64(i)))
			found := false
			for j := 3; j <= l; j += 2 {
				if i%j == 0 {
					found = true
					break
				}
			}
			if !found {
				result <- i // 送信!
			}
		}
		close(result) // 終了情報のシグナルを目的としたチャネルは明示的に close() しよう
	}()

	return result
}

func main() {
	pn := primeNumber()

	for n := range pn { // 送信! のたびにループが回る
		fmt.Println(n)
	}
}
