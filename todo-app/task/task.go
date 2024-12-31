package task

import (
	"fmt"
	"github.com/mkashifaslam/golang/todo-app/store"
	"github.com/mkashifaslam/golang/todo-app/utils"
)

var tasks []Task

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
	return store.SaveToFile(t)
}

func (t *Task) Complete() {
	t.IsCompleted = true
}

func (t *Task) Print() {
	fmt.Printf("ID: %d, Title: %s, IsComplete: %t\n", t.ID, t.Title, t.IsCompleted)
}

func AddToList(t *Task) {
	tasks = append(tasks, *t)
}

func PrintList() {
	for _, t := range tasks {
		t.Print()
	}
}

func PrintTasks() error {
	lines, err := store.ReadFromFile()
	if err != nil {
		return err
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	return nil
}
