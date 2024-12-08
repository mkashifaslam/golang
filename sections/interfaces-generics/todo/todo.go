package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Content string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Println(todo.Content)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	jsonValue, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonValue, 0644)
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("invalid input")
	}

	return Todo{
		Content: content,
	}, nil
}
