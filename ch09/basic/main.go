package main

import (
	"fmt"
	"io"
	"os"
	"syscall"
	"time"
)

// 新規作成
func create(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.WriteString(file, "New file contents\n")
}

// 読み込み
func open(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println("Read file:")
	io.Copy(os.Stdout, file)
}

// 追記
func appendFile(filename string) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.WriteString(file, "Appended content\n")
}

func main() {
	fmt.Println("Basic Go standard library about file system")
	fmt.Println("--------------------------------------------")

	filename := "testfile.txt"

	create(filename)
	open(filename)

	appendFile(filename)

	fmt.Println("--------------------------------------------")

	// ディレクトリの作成
	os.Mkdir("setting", 0755)
	// 深いディレクトリを1回で作成
	os.MkdirAll("setting/myapp/networksettings", 0755)
	// owner, group, others
	// 0755 -> rwxr-xr-x / 0644 -> rw-r--r--

	// ファイルや空のディレクトリを削除
	os.Remove("server.log")
	// ディレクトリを中身ごと削除
	os.RemoveAll("workdir")

	// 先頭100バイトで切り落とす
	os.Truncate("server.log", 100)
	// os.File に対して Truncate メソッドを適用する場合
	file, _ := os.Open("server.log")
	file.Truncate(100)

	// リネーム
	os.Rename("old_name.txt", "new_name.txt")
	// 移動
	os.Rename("olddir/file.txt", "newdir/file.txt")
	// 移動先はディレクトリだけでなくファイル名まで指定する
	// os.Rename("olddir/file.txt", "newdir/") // -> error

	// デバイスやドライバが異なる先へのファイルコピーは、ファイルを開いて中身をコピーする必要がある

	// ファイルの情報確認
	if len(os.Args) == 1 {
		fmt.Printf("%s [exec file name]\n", os.Args[0])
		os.Exit(1)
	}
	info, err := os.Stat(os.Args[1])
	if err == os.ErrNotExist {
		fmt.Printf("file not found: %s\n", os.Args[1])
	} else if err != nil {
		panic(err)
	}
	fmt.Println("FileInfo")
	fmt.Printf("  ファイル名: %v\n", info.Name())
	fmt.Printf("  サイズ名: %v\n", info.Size())
	fmt.Printf("  変更日時: %v\n", info.ModTime())
	fmt.Println("Mode()")
	fmt.Printf("  ディレクトリ？: %v\n", info.Mode().IsDir())
	fmt.Printf("  読み書き可能な通常ファイル？: %v\n", info.Mode().IsRegular())
	fmt.Printf("  Unixのファイルアクセス権限ビット: %o\n", info.Mode().Perm())
	fmt.Printf("  モードのテキスト表現: %v\n", info.Mode().String())

	fmt.Println("--------------------------------------------")

	// OS固有のファイル属性を取得 -> OS固有の構造体にFileInfo.Sys()をダウンキャスト
	// Windows
	// internalStat := info.Sys().(syscall.Win32FileAttributeData)
	// Others
	internalStat := info.Sys().(*syscall.Stat_t)
	fmt.Println("internalStat")
	fmt.Printf("  デバイス番号: %v\n", internalStat.Dev)
	fmt.Printf("  inode番号: %v\n", internalStat.Ino)
	fmt.Printf("  ブロックサイズ: %v\n", internalStat.Blksize)
	fmt.Printf("  ブロック数: %v\n", internalStat.Blocks)
	fmt.Printf("  リンクされている数: %v\n", internalStat.Nlink)
	fmt.Printf("  最終アクセス日時: %v\n", internalStat.Atimespec)
	fmt.Printf("  属性変更日時: %v\n", internalStat.Ctimespec)

	fmt.Println("--------------------------------------------")

	// ファイルの同一性チェック（同値ではなく同じ実体かどうかを判定）
	info2, _ := os.Stat(os.Args[1])
	if os.SameFile(info, info2) {
		fmt.Println("same file")
	}

	// ファイルの属性の設定
	os.Chmod("setting.txt", 0644)
	// ファイルのオーナーを変更
	os.Chown("setting.txt", os.Getuid(), os.Getgid())
	// ファイルの最終アクセス日時と変更日時を変更
	os.Chtimes("setting.txt", time.Now(), time.Now())

	// ハードリンクの作成
	os.Link("oldfile.txt", "newfile.txt")
	// シンボリックリンクの作成
	os.Symlink("oldfile.txt", "newfile-symlink.txt")
	// シンボリックリンクのリンク先を取得
	_, err = os.Readlink("newfile-symlink.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("--------------------------------------------")

	// ディレクトリ情報の取得
	dir, err := os.Open("/")
	if err != nil {
		panic(err)
	}
	fileInfos, err := dir.Readdir(-1) // ディレクトリ内の全要素
	if err != nil {
		panic(err)
	}
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			fmt.Printf("[Dir]  %s\n", fileInfo.Name())
		} else {
			fmt.Printf("[File] %s\n", fileInfo.Name())
		}
	}
}
