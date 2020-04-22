package main

import (
	crand "crypto/rand"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	mrand "math/rand"
	"time"
)

func main() {
	a := make([]byte, 20)
	crand.Read(a)
	fmt.Println(hex.EncodeToString(a))

	mrand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		fmt.Println(mrand.Float64())
	}

	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	src := mrand.NewSource(seed.Int64())
	rng := mrand.New(src) // 乱数生成器のオブジェクトを必要な数だけ作成->ユーザーごとに独立させることが出来る
	fmt.Println(rng.Uint64())
}
