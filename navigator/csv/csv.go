package csv

import (
	"bufio"
	"bytes"
	"embed"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"strings"
)

type Data map[string]map[string]string

func LoadData(fileName string, fs embed.FS) (Data, error) {
	log.SetFlags(0)
	//log.Println("filename", fileName)
	if fileName == "" {
		return nil, errors.New("file name is required")
	}

	//records, err := readOSFile(fileName)
	records, err := readFile(fileName, fs)
	if err != nil {
		//fmt.Println("Error reading CSV file:", err)
		return nil, errors.New("Error reading CSV file: " + err.Error())
	}
	//log.Println("file content checking...")
	// Check if the CSV is not empty
	if len(records) < 2 {
		//fmt.Println("CSV file is empty or has no data rows.")
		return nil, errors.New("csv file is empty or has no data rows")
	}

	//log.Println("file content exists...")
	// Get the header (column names)
	header := records[0]

	// Create a map to store the CSV data
	dataMap := make(map[string]map[string]string)

	// Iterate through the rows and populate the map
	for _, row := range records[1:] {
		// Assuming the first column is the key
		key := row[0]

		// Create a nested map for each row
		rowMap := make(map[string]string)
		for i, value := range row[1:] {
			rowMap[header[i+1]] = value
		}

		// Add the row map to the main map
		dataMap[key] = rowMap
	}

	return dataMap, nil
}

func readFile(fileName string, fs embed.FS) ([][]string, error) {
	//file, err := os.Open(fileName)
	file, err := fs.Open(fileName)
	//log.Println("file open...")
	if err != nil {
		return nil, errors.New("Error opening file: " + err.Error())
	}

	defer file.Close()

	//log.Println("file open success")
	// Read the file using csv.Reader
	reader := csv.NewReader(file)

	// Read all rows (including the header)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, errors.New("Error reading file: " + err.Error())
	}

	return records, nil
}

func scanFile(fileName string, fs embed.FS) ([][]string, error) {
	content, err := fs.ReadFile(fileName)
	if err != nil {
		return nil, errors.New("Error reading file: " + err.Error())
	}

	var records [][]string
	scanner := bufio.NewScanner(bytes.NewReader(content))
	for scanner.Scan() {
		records = append(records, strings.Split(scanner.Text(), ","))
	}

	return records, nil
}

func Print(data Data) {
	// Print the map
	for key, value := range data {
		fmt.Printf("Key: %s, Value: %v\n", key, value)
	}
}
