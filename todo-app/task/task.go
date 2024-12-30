package task

import "fmt"

var tasks []Task

type Task struct {
	Title       string
	IsCompleted bool
}

func New(title string) *Task {
	return &Task{
		Title: title,
	}
}

func (t *Task) Complete() {
	t.IsCompleted = true
}

func (t *Task) Print() {
	fmt.Printf("Title: %s, IsComplete: %t\n", t.Title, t.IsCompleted)
}

func AddToList(t *Task) {
	tasks = append(tasks, *t)
}

func PrintList() {
	for _, t := range tasks {
		t.Print()
	}
}
