package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

// 取得してきたサーバーの鍵情報
var hostKeyString string = "ecdsa-sha2-nistp256 AAAAE...idDI="

func main() {
	// 秘密鍵の準備
	key, err := ioutil.ReadFile("id_sysprogo")
	if err != nil {
		panic(err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		panic(err)
	}

	hostKey, _, _, _, err := ssh.ParseAuthorizedKey([]byte(hostKeyString))
	if err != nil {
		panic(err)
	}

	// 接続設定
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(singer), // 公開鍵方式
			// ssh.Password("root-password"), // パスワード方式
		},
		Timeout:         5 * time.Second,
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	// 通信開始
	conn, err := ssh.Dial("tcp", "localhost:1222", config)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// コマンドを実行して出力結果を取得
	output, err := session.CombinedOutput("ssh -v")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))

	// シェルモード
	// session.Stdout = os.Stdout
	// session.Stderr = os.Stderr
	// session.Stdin = os.Stdin
	// session.Shell()
	// session.Wait()
}
