package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("Go path/filepath")
	fmt.Println("----------------------------------------")

	// filepath.Join()
	fmt.Printf("Temp File Path: %s\n", filepath.Join(os.TempDir(), "temp.txt"))

	fmt.Println("----------------------------------------")

	// filepath.Split()
	dir, name := filepath.Split(os.Getenv("GOPATH"))
	fmt.Printf("Dir: %s, Name: %s\n", dir, name)

	fmt.Println("----------------------------------------")

	// filepath.SplitList() で which コマンド実装
	if len(os.Args) == 1 {
		fmt.Printf("%s [exec file name]\n", os.Args[0])
		os.Exit(1)
	}
	for _, path := range filepath.SplitList(os.Getenv("PATH")) {
		execpath := filepath.Join(path, os.Args[1])
		_, err := os.Stat(execpath)
		if !os.IsNotExist(err) {
			fmt.Println(execpath)
		}
	}

	fmt.Println("----------------------------------------")

	// pathのクリーン
	fmt.Println(filepath.Clean("./path/filepath/../path.go"))
	// パスを絶対パスに
	abspath, _ := filepath.Abs("path/filepath/path_unix.go")
	fmt.Println(abspath)
	// パスを相対パスに
	relpath, _ := filepath.Rel("usr/local/go/src", "usr/local/go/src/path/filepath/path.go")
	fmt.Println(relpath)

	fmt.Println("----------------------------------------")

	// 環境変数の展開 -> path/filepath ではなくて os
	path := os.ExpandEnv("${GOPATH}/src/github.com/matsuyoshi30")
	fmt.Println(path)
	// ~ -> OSではなくてシェルが提供する機能なので少し工夫が必要
	Clean2 := func(path string) string {
		if len(path) > 1 && path[0:2] == "~/" { // ~を置換
			my, err := user.Current()
			if err != nil {
				panic(err)
			}
			path = my.HomeDir + path[1:]
		}
		return filepath.Clean(path)
	}
	fmt.Println(Clean2("~/Downloads"))

	fmt.Println("----------------------------------------")

	// ファイル名のパターンマッチ * ? [0-9]
	fmt.Println(filepath.Match("image-*.png", "image-100.png"))
	// マッチするファイル名の一覧を取得
	files, err := filepath.Glob("./*.go")
	if err != nil {
		panic(err)
	}
	fmt.Println(files)

	fmt.Println("----------------------------------------")

	// ディレクトリのトラバース
	// filepath.Walk() が探索を開始するノードと探索過程で処理されるコールバック関数を引数にとり、深さ優先探索でディレクトリをトラバースする
	var imageSuffix = map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".webp": true,
		".gif":  true,
		".tiff": true,
		".eps":  true,
	}
	if len(os.Args) == 1 {
		fmt.Printf(`Find images

usage:
    %s [path to find]
`, os.Args[0])
		return
	}
	root := os.Args[2]

	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if info.Name() == "__build" { // 特定のディレクトリ配下はスキップ
				return filepath.SkipDir
			}
			return nil // それ以外のディレクトリは、ディレクトリ自体は処理せずファイルにあたるまで再帰的に処理
		}
		ext := strings.ToLower(filepath.Ext(info.Name())) // .PNG とかに対応
		if imageSuffix[ext] {
			rel, err := filepath.Rel(root, path) // pathを整形
			if err != nil {
				return nil
			}
			fmt.Printf("%s\n", rel)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
