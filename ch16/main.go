package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fs := 5 * time.Second
	tms := 10 * time.Millisecond
	tmts, _ := time.ParseDuration("10m30s")
	fmt.Println(fs, tms, tmts)

	fmt.Println(time.Now())
	fmt.Println(time.Date(2017, time.August, 26, 11, 50, 30, 0, time.Local))
	fmt.Println(time.Parse(time.Kitchen, "11:30AM"))
	fmt.Println(time.Unix(1503673200, 0))

	fmt.Println(time.Now().Add(3 * time.Hour))
	fileInfo, _ := os.Stat("main.go")
	fmt.Printf("%v 前\n", time.Now().Sub(fileInfo.ModTime()))
	fmt.Println(time.Now().Round(time.Hour))

	fmt.Println("waiting 3 seconds")
	time.Sleep(3 * time.Second)
	fmt.Println("done")

	fmt.Println("waiting 3 seconds by channel")
	after := time.After(3 * time.Second)
	<-after
	fmt.Println("done")

	fmt.Println(time.Now().Format(time.RFC822))
	fmt.Println(time.Now().Format("2006/01/02 03:04:05 MST"))

	fmt.Println("waiting 3 seconds by tick")
	for now := range time.Tick(3 * time.Second) {
		fmt.Println("now:", now)
	}
}
