package main

import (
	"fmt"
	"time"
)

func sub1(c int) {
	fmt.Println("share by arguments:", c)
}

func sub2(s string) {
	fmt.Println(s)
}

func main() {
	// 引数渡し
	go sub1(10)

	// クロージャのキャプチャ渡し
	c := 10
	go func() {
		fmt.Println("share by capture:", c*c)
	}()
	time.Sleep(time.Second) // これないとmainが先に終了してしまうので説明用

	// クロージャのキャプチャ渡しは内部的には無名関数に暗黙の引数が追加され、
	// 暗黙の引数にデータや参照が渡されてgoroutineとして扱われる
	// -> 上の２つは結果的に同じ
	// https://speakerdeck.com/rhysd/go-detukurufan-yong-yan-yu-chu-li-xi-shi-zhuang-zhan-lue

	// 引数渡しとクロージャのキャプチャ渡しの違い
	tasks := []string{
		"cmake ..",
		"cmake . --build Release",
		"cpack",
	}
	for _, task := range tasks {
		go func() {
			// goroutineを起動するときにはループが回りきって
			// 全taskが最後のタスクになってしまう
			fmt.Println(task)
		}()
	}
	for _, task := range tasks {
		// 全種類のtaskが出力される
		go sub2(task)
	}
	time.Sleep(time.Second)
}
