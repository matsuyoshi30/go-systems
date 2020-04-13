package main

import (
	"io"
	"log"
	"os"
	"strings"
)

// io.CopyN(dest io.Writer, src io.Reader, length int))

func main() {
	f, err := os.Create("dest.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := strings.NewReader("This is another io.CopyN.")
	// io.CopyN(f, r, 8)
	io.Copy(f, io.LimitReader(r, 8))
}
