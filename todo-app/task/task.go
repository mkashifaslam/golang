package task

import (
	"encoding/json"
	"fmt"
	"github.com/mkashifaslam/golang/todo-app/store"
	"github.com/mkashifaslam/golang/todo-app/utils"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"is_completed"`
}

func New(title string) *Task {
	return &Task{
		ID:    utils.GenerateRandomInt(1, 100),
		Title: title,
	}
}

func (t *Task) Save() error {
	return store.Append(t)
}

func (t *Task) Complete() {
	t.IsCompleted = true
}

func (t *Task) Delete() {
}

func (t *Task) Print() {
	fmt.Printf("ID: %d, Title: %s, IsComplete: %t\n", t.ID, t.Title, t.IsCompleted)
}

func GetTasks() ([]Task, error) {
	lines, err := store.Read()

	if err != nil {
		return nil, err
	}

	var tasks []Task
	for _, line := range lines {
		var task Task
		_ = json.Unmarshal([]byte(line), &task)
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func PrintTasks() error {
	lines, err := store.Read()
	if err != nil {
		return err
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	return nil
}

func GetTaskByID(id int) (*Task, error) {
	var tasks, err = GetTasks()

	if err != nil {
		return nil, utils.ErrorHandler(err, "TasksLoadingFailed:")
	}

	var task Task
	for _, t := range tasks {
		if t.ID == id {
			task = t
			break
		}
	}

	if task.ID == 0 {
		return nil, utils.NewError("task not found")
	}

	return &task, nil
}
