package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"

	"github.com/shirou/gopsutil/process"

	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"

	"github.com/kr/pty"
)

func main() {
	// プロセスに含まれている情報の確認
	fmt.Println("process id:", os.Getpid())
	fmt.Println("parent process id:", os.Getppid())

	sid, _ := syscall.Getsid(os.Getpid())
	fmt.Fprintf(os.Stderr, "process group id: %d, session id: %d\n", syscall.Getpgrp(), sid)

	fmt.Printf("user id: %d\n", os.Getuid())
	fmt.Printf("group id: %d\n", os.Getgid())
	groups, _ := os.Getgroups()
	fmt.Printf("sub group id: %v\n", groups)

	fmt.Printf("user id: %d\n", os.Getuid())
	fmt.Printf("group id: %d\n", os.Getgid())
	fmt.Printf("effective user id: %d\n", os.Geteuid())
	fmt.Printf("effective group id: %d\n", os.Getegid())

	wd, _ := os.Getwd()
	fmt.Println(wd)

	// 環境変数
	fmt.Println(os.ExpandEnv("${HOME}/gobin"))

	// 終了コード
	// 0 -> 正常終了, 1以上 -> エラー終了
	//os.Exit(1)

	// プロセスIDが何者なのか
	p, _ := process.NewProcess(int32(os.Getppid()))
	name, _ := p.Name()
	cmd, _ := p.Cmdline()
	memPer, _ := p.MemoryPercent()
	cpuPer, _ := p.CPUPercent()
	fmt.Printf("parent pid: %d, name: '%s', cmd: '%s'\n", p.Pid, name, cmd)
	fmt.Printf("memory percent: %f, cpu percent: %f\n", memPer, cpuPer)

	// exec.Command()の例
	if len(os.Args) == 1 {
	} else {
		command := exec.Command(os.Args[1], os.Args[2:]...) // 引数に外部コマンドを取る
		err := command.Run()
		if err != nil {
			panic(err)
		}
		state := command.ProcessState
		fmt.Printf("%s\n", state.String())
		fmt.Printf("  Pid: %d\n", state.Pid())
		fmt.Printf("  System: %v\n", state.SystemTime()) // カーネル内で行われた処理の時間
		fmt.Printf("  User: %v\n", state.UserTime())     // プロセス内で消費された時間
	}

	// exec.Cmd()の子プロセスとのリアルタイム通信
	count := exec.Command("./count")
	stdout, _ := count.StdoutPipe()
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			fmt.Printf("(stdout) %s\n", scanner.Text())
		}
	}()
	err := count.Run()
	if err != nil {
		panic(err)
	}

	// mattn library
	var data = "\033[34m\033[47m\033[4mB\033[31me\n\033[24m\033[30mOS\033[49m\033[m\n"
	var stdOut io.Writer
	if isatty.IsTerminal(os.Stdout.Fd()) {
		stdOut = colorable.NewColorableStdout()
	} else { // ファイルにリダイレクトするとエスケープシーケンスが出力されないことが確認できる！
		stdOut = colorable.NewNonColorable(os.Stdout)
	}
	fmt.Fprintln(stdOut, data)

	// 擬似端末詐称
	checkCmd := exec.Command("./check")
	stdpty, stdtty, _ := pty.Open()
	defer stdtty.Close()
	checkCmd.Stdin = stdpty
	checkCmd.Stdout = stdpty
	errpty, errtty, _ := pty.Open()
	defer errtty.Close()
	checkCmd.Stderr = errtty
	go func() {
		io.Copy(os.Stdout, stdpty)
	}()
	go func() {
		io.Copy(os.Stderr, errpty)
	}()
	err = checkCmd.Run()
	if err != nil {
		panic(err)
	}
}
