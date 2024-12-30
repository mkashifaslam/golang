package app

import (
	"fmt"
	"github.com/mkashifaslam/golang/todo-app/input"
	t "github.com/mkashifaslam/golang/todo-app/task"
)

func Run() {
	action, err := input.GetUserInput()
	if err != nil {
		fmt.Println("Error reading user input", err.Error())
	}

	title := input.GetUserStringInput("Title")
	execAction(action, title)

	t.PrintList()
	fmt.Println("Done")
}

func execAction(action string, title string) {
	switch action {
	case "add":
		fmt.Println("Adding task...", title)
		task := t.New(title)
		t.AddToList(task)
	case "complete":
		fmt.Println("Complete task")
	default:
		fmt.Println("Unknown action")
	}
}
