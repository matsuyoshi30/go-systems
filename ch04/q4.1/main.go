package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("set timer value")
	}

	t, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	ch := time.After(time.Duration(t) * time.Second)
	fmt.Println("start")
	<-ch
	fmt.Println("end")
}
