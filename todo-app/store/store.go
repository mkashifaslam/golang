package store

import (
	"bufio"
	"encoding/json"
	"github.com/mkashifaslam/golang/todo-app/utils"
	"os"
)

const OutputFileName = "todos.txt"

func Read() ([]string, error) {
	file, err := os.OpenFile(OutputFileName, os.O_RDONLY, 0644)

	if err != nil {
		return nil, utils.ErrorHandler(err, "failed to open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		//_ = file.Close()
		return nil, utils.ErrorHandler(err, "failed to read file content")
	}

	return lines, nil
}

func Append(data any) error {
	file, err := os.OpenFile(OutputFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		return utils.ErrorHandler(err, "failed to create file")
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return utils.ErrorHandler(err, "failed to convert data to json")
	}

	return nil
}

func Write(data any) error {
	file, err := os.OpenFile(OutputFileName, os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		return utils.ErrorHandler(err, "failed to write file")
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return utils.ErrorHandler(err, "failed to convert data to json")
	}

	return nil
}