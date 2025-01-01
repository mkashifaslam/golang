package store

import (
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

func Write(data string) error {
	err := os.WriteFile(OutputFileName, []byte(data), 0644)

	if err != nil {
		return utils.ErrorHandler(err, "failed to write file")
	}

	return nil
}
