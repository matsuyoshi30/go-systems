package main

import (
	"fmt"
)

type S struct {
	name string
}

func main() {
	// プリミティブのインスタンスを定義
	var a int = 10
	// 構造体のインスタンスをnewして作成
	// 変数にはポインタを保存
	var b *S = new(S)
	// 構造体を{}でメンバーの初期値を与えて初期化
	// 変数にはインスタンスを保存
	var c S = S{"param"}
	// 構造体を{}でメンバーの初期値を与えて初期化
	// 変数にはポインタを保存
	var d *S = &S{"param"}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	// 固定長配列
	e := [4]int{1, 2, 3, 4}
	// len, capを持ったスライス
	f := make([]int, 4, 8)
	// バッファなしのチャネル
	g := make(chan int)
	// バッファありのチャネル
	h := make(chan string, 10)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)

	// 既存の配列を参照するスライス
	x := [4]int{1, 2, 3, 4}
	y := x[:]
	fmt.Println(&y[0], len(y), cap(y))
	// 既存の配列の一部を参照するスライス
	z := x[1:3]
	fmt.Println(&z[0], len(z), cap(z)) // -> address, 2, 3

	// 何も参照していないスライス
	var w []int
	fmt.Println(len(w), cap(w))
}
