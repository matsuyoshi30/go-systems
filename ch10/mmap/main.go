package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/edsrzf/mmap-go"
)

func main() {
	var testData = []byte("0123456789ABCDEF")
	var testPath = filepath.Join(os.TempDir(), "testdata")
	err := ioutil.WriteFile(testPath, testData, 0644)
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile(testPath, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	m, err := mmap.Map(f, mmap.RDWR, 0) // 指定したファイルをメモリ上に展開
	// m, err := mmap.Map(f, mmap.COPY, 0) // コピーオンライト
	if err != nil {
		panic(err)
	}
	defer m.Unmap() // メモリ上に展開された内容を削除して閉じる

	m[9] = 'X'
	m.Flush() // 書きかけの内容をファイルに保存する

	fileData, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("original: %s\n", testData)
	fmt.Printf("mmap:     %s\n", m)
	fmt.Printf("file:     %s\n", fileData)
}
