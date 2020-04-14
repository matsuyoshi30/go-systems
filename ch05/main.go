package main

import (
	"fmt"
	"log"
	_ "net/http/pprof"
	"os"
	"runtime/trace"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	trace.Start(f)
	defer trace.Stop()

	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for i := 1; i <= 100; i++ {
		fmt.Fprintf(file, "Line %d\n", i)
	}
}
