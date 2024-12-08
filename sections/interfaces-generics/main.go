package main

import (
	"bufio"
	"fmt"
	"github.com/golang/interfaces-generics/note"
	"github.com/golang/interfaces-generics/todo"
	"os"
	"strings"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	todoText := getTodoData()
	userTodo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userTodo.Display()
	err = userTodo.Save()

	if err != nil {
		fmt.Println("Saving the todo failed.")
		return
	}

	fmt.Println("Saving the note succeeded!")

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}

	fmt.Println("Saving the note succeeded!")
}

func getTodoData() string {
	return getUserInput("Todo")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content

}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
