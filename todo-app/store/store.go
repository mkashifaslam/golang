package store

import (
	"encoding/json"
	"github.com/mkashifaslam/golang/todo-app/utils"
	"os"
)

const OutputFileName = "todos.txt"

func Read() (string, error) {
	content, err := os.ReadFile(OutputFileName)
	if err != nil {
		return "", utils.ErrorHandler(err, "failed to read file")
	}

	return string(content), nil
}

func Write(data any) error {
	file, err := os.Create(OutputFileName)

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
