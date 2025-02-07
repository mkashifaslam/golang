package csv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

type Data map[string]map[string]string

func LoadData(fileName string) (Data, error) {

	if fileName == "" {
		return nil, errors.New("file name is required")
	}

	file, err := os.Open(fileName)
	if err != nil {
		//fmt.Println("Error opening file:", err)
		return nil, errors.New("Error opening file: " + err.Error())
	}
	defer file.Close()

	// Read the file using csv.Reader
	reader := csv.NewReader(file)

	// Read all rows (including the header)
	records, err := reader.ReadAll()
	if err != nil {
		//fmt.Println("Error reading CSV file:", err)
		return nil, errors.New("Error reading CSV file: " + err.Error())
	}

	// Check if the CSV is not empty
	if len(records) < 2 {
		//fmt.Println("CSV file is empty or has no data rows.")
		return nil, errors.New("csv file is empty or has no data rows")
	}

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

func Print(data Data) {
	// Print the map
	for key, value := range data {
		fmt.Printf("Key: %s, Value: %v\n", key, value)
	}
}
