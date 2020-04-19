package main

import (
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Accept Ctrl+C for 5 seconds")
	time.Sleep(time.Second * 5)

	signal.Ignore(syscall.SIGINT, syscall.SIGHUP)

	fmt.Println("Ignore Ctrl+C for 5 seconds")
	time.Sleep(time.Second * 5)
}
