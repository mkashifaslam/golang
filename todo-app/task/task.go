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
	tasks, err := GetTasks()
	if err != nil {
		utils.PrintError(err, "TasksLoadingFailed:")
	}

	tasks = append(tasks, *t)

	tasksToJson, err := TasksToJson(tasks)
	if err != nil {
		utils.PrintError(err, "TasksToJsonError:")
	}

	return store.Write(tasksToJson)
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
	content, err := store.Read()

	if err != nil {
		return nil, err
	}

	tasks, err := JsonToTasks(content)

	return tasks, nil
}

func PrintTasks() error {
	content, err := store.Read()
	if err != nil {
		return err
	}

	tasks, err := JsonToTasks(content)

	for _, task := range tasks {
		task.Print()
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

func CompleteTaskById(tasks []Task, id int) ([]Task, error) {
	for i, t := range tasks {
		if t.ID == id {
			t.Complete()
			tasks[i] = t
			return tasks, nil
		}
	}

	return nil, utils.NewError("task not found")
}

func DeleteTaskByID(tasks []Task, id int) ([]Task, error) {
	for i, t := range tasks {
		if t.ID == id {
			return append(tasks[:i], tasks[i+1:]...), nil
		}
	}

	return nil, utils.NewError("task not found")
}

func JsonToTasks(content string) ([]Task, error) {
	var tasks []Task
	err := json.Unmarshal([]byte(content), &tasks)
	return tasks, err
}

func TasksToJson(tasks []Task) (string, error) {
	content, err := json.Marshal(tasks)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
