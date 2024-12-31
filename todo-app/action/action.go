package action

import (
	"fmt"
	"github.com/mkashifaslam/golang/todo-app/input"
	t "github.com/mkashifaslam/golang/todo-app/task"
	"github.com/mkashifaslam/golang/todo-app/utils"
	"strconv"
)

func GetActionInput() string {
	userInput, err := input.GetUserInput()
	if err != nil {
		utils.PrintError(err, "Error reading user input")
	}
	return userInput
}

func Run(action string) {
	switch action {
	case "add":
		addTask()
	case "list":
		listTasks()
	case "find":
		findTask()
	case "complete":
		completeTask()
	default:
		fmt.Println("Unknown action")
	}
}

func addTask() {
	title := input.GetUserStringInput("Title")
	task := t.New(title)
	err := task.Save()
	if err != nil {
		utils.PrintError(err, "Error saving task")
	}
}

func listTasks() {
	err := t.PrintTasks()
	if err != nil {
		utils.PrintError(err, "Error listing tasks")
	}
}

func findTaskById() (*t.Task, error) {
	taskId := input.GetUserStringInput("TaskId")
	taskIdInt64, err := strconv.ParseInt(taskId, 10, 64)
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error parsing taskId")
	}
	return t.GetTaskByID(int(taskIdInt64))
}

func findTask() *t.Task {
	task, err := findTaskById()

	if err != nil {
		utils.PrintError(err, "Error finding task")
	}

	task.Print()

	return task
}

func completeTask() {
	task := findTask()
	task.Complete()
}
