package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	f, err := os.Create("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	writer1 := csv.NewWriter(f)         // output to file
	writer2 := csv.NewWriter(os.Stdout) // output to stdout
	for _, record := range records {
		if err = writer1.Write(record); err != nil {
			log.Fatal(err)
		}
		if err = writer2.Write(record); err != nil {
			log.Fatal(err)
		}
	}
	writer1.Flush()
	writer2.Flush()
}
