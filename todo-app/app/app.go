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

	execAction(action)

	fmt.Println("Done")
}

func execAction(action string) {
	switch action {
	case "add":
		title := input.GetUserStringInput("Title")
		task := t.New(title)
		_ = task.Save()
	case "list":
		_ = t.PrintTasks()
	case "find":
		fmt.Println("task found")
	case "complete":
		fmt.Println("Complete task")
	default:
		fmt.Println("Unknown action")
	}
}
