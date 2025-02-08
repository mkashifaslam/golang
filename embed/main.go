package main

import (
	"embed"
	"encoding/csv"
	"fmt"
)

//go:embed data/cities.csv
var fs embed.FS

func main() {
	file, err := fs.Open("data/cities.csv")
	//content, err := fs.ReadFile("data/cities.csv")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer file.Close()

	//log.Println("file open success")
	// Read the file using csv.Reader
	reader := csv.NewReader(file)

	// Read all rows (including the header)
	records, err := reader.ReadAll()
	fmt.Println(records)
}
