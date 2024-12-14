package prices_list

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

func GetPricesFromFile(fileName string) ([]int, error) {
	prices, err := os.ReadFile(fileName)
	if err != nil {
		return nil, errors.New("Failed to find prices file " + fileName)
	}

	return parsePricesList(prices)
}

func parsePricesList(data []byte) ([]int, error) {
	// Split the file content into lines
	lines := strings.Split(string(data), "\n")

	// Slice to store the numbers
	var numbers []int

	// Convert each line to an integer
	for _, line := range lines {
		// Trim any whitespace
		line = strings.TrimSpace(line)
		if line == "" { // Skip empty lines
			continue
		}

		// Convert to integer
		number, err := strconv.Atoi(line)
		if err != nil {
			return nil, errors.New("Error" + err.Error())
		}

		// Append to the slice
		numbers = append(numbers, number)
	}

	return numbers, nil

}
