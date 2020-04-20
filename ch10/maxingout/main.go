package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

// 簡単にまとめるとこんな感じでエラー
// https://play.golang.org/p/k3SDPbvDPeb

func main() {
	cases := [65537]reflect.SelectCase{}
	for i, _ := range cases {
		cases[i] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(make(chan int)),
		}
	}

	go func() {
		for {
			i := rand.Int() % len(cases)
			cases[i].Chan.Send(reflect.ValueOf(i))
			time.Sleep(time.Millisecond)
		}
	}()

	fmt.Println("ready")

	for {
		c, v, ok := reflect.Select(cases[:])
		fmt.Println(c, v, ok)
	}
}
