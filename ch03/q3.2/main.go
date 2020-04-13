package main

import (
	"crypto/rand"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Create("randfile")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err = io.CopyN(f, rand.Reader, 1024); err != nil {
		log.Fatal(err)
	}
}
