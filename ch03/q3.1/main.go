package main

import (
	"io"
	"log"
	"os"
)

func main() {
	oldf, err := os.Open("old.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer oldf.Close()

	newf, err := os.Create("new.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newf.Close()

	io.Copy(newf, oldf)
}
