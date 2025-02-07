package action

import (
	"fmt"
	"github.com/mkashifaslam/golang/todo-app/store"
	t "github.com/mkashifaslam/golang/todo-app/task"
	"github.com/mkashifaslam/golang/todo-app/utils"
	"strconv"
)

type Act string

const (
	Add      Act = "add"
	List     Act = "list"
	Find     Act = "find"
	Delete   Act = "delete"
	Complete Act = "complete"
)

func Run(action Act, args string) {
	switch action {
	case Add:
		addTask(args)
	case List:
		listTasks()
	case Find:
		findTask(args)
	case Complete:
		completeTask(args)
	case Delete:
		deleteTask(args)
	default:
		fmt.Println("Unknown action")
	}
}

func addTask(title string) {
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

func findTaskById(taskId string) (*t.Task, error) {
	taskIdInt64, err := strconv.ParseInt(taskId, 10, 64)

	if err != nil {
		return nil, utils.ErrorHandler(err, "ParsingError:")
	}

	return t.GetTaskByID(int(taskIdInt64))
}

func findTask(taskId string) *t.Task {
	task, err := findTaskById(taskId)

	if err != nil {
		utils.PrintError(err, "FindError:")
		return nil
	}

	task.Print()

	return task
}

func completeTask(taskId string) {
	taskIdInt64, err := strconv.ParseInt(taskId, 10, 64)

	if err != nil {
		utils.PrintError(err, "ParsingError:")
	}

	tasks, err := t.GetTasks()
	if err != nil {
		utils.PrintError(err, "TasksLoadingError:")
	}

	tasks, err = t.CompleteTaskById(tasks, int(taskIdInt64))
	if err != nil {
		utils.PrintError(err, "TaskCompletedError:")
	}

	tasksToJson, err := t.TasksToJson(tasks)
	if err != nil {
		utils.PrintError(err, "TasksToJsonError:")
	}

	err = store.Write(tasksToJson)
	if err != nil {
		utils.PrintError(err, "TasksUpdateError:")
	}
}

func deleteTask(taskId string) {
	taskIdInt64, err := strconv.ParseInt(taskId, 10, 64)

	if err != nil {
		utils.PrintError(err, "ParsingError:")
	}

	tasks, err := t.GetTasks()
	if err != nil {
		utils.PrintError(err, "TasksLoadingError:")
	}

	tasks, err = t.DeleteTaskByID(tasks, int(taskIdInt64))
	if err != nil {
		utils.PrintError(err, "TaskDeleteError:")
	}

	tasksToJson, err := t.TasksToJson(tasks)
	if err != nil {
		utils.PrintError(err, "TasksToJsonError:")
	}

	err = store.Write(tasksToJson)
	if err != nil {
		utils.PrintError(err, "TasksUpdateError:")
	}

}
