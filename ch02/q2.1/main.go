package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("tempfile")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Fprintf(f, "number %d\n", 10)
	fmt.Fprintf(f, "string %s\n", "str")
	fmt.Fprintf(f, "float  %f\n", 4.4)
}
