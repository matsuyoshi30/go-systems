package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	dist, err := os.Create("distfile.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer dist.Close()

	zipwriter := zip.NewWriter(dist)
	defer zipwriter.Close()

	w, err := zipwriter.Create("newfile")
	if err != nil {
		log.Fatal(err)
	}

	r := strings.NewReader("this is test string.")
	io.Copy(w, r)
}
