package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	source := map[string]string{
		"Hello": "world",
	}

	gw := gzip.NewWriter(w)
	writer := io.MultiWriter(gw, os.Stdout)

	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "    ")
	encoder.Encode(source)

	gw.Flush()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
