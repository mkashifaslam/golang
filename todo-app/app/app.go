package app

import (
	"fmt"
	"github.com/mkashifaslam/golang/todo-app/input"
	t "github.com/mkashifaslam/golang/todo-app/task"
	"github.com/mkashifaslam/golang/todo-app/utils"
)

func Run() {
	action, err := input.GetUserInput()
	if err != nil {
		utils.PrintError(err, "Error reading user input")
	}

	title := input.GetUserStringInput("Title")
	task := t.New(title)
	execAction(action, task)

	t.PrintList()
	fmt.Println("Done")
}

func execAction(action string, task *t.Task) {
	switch action {
	case "add":
		t.AddToList(task)
	case "complete":
		fmt.Println("Complete task")
	default:
		fmt.Println("Unknown action")
	}
}
