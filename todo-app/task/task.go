package task

import (
	"fmt"
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
