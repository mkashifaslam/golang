package main

import (
	"bufio"
	"fmt"
	"github.com/golang/interfaces-generics/note"
	"github.com/golang/interfaces-generics/todo"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

type outputter interface {
	saver
	Display()
}

func main() {

	//printSomething(1)
	//printSomething(2.5)
	//printSomething("hello!")
	//
	//intResult := add(1, 2)
	//fmt.Println(intResult)
	//
	//floatResult := add(1.3, 2.6)
	//fmt.Println(floatResult)
	//
	//strResult := add("hello,", "world")
	//fmt.Println(strResult)

	title, content := getNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	todoText := getTodoData()
	userTodo, err := todo.New(todoText)

	//printSomething(userTodo)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userTodo)

	if err != nil {
		return
	}

	err = outputData(userNote)

	if err != nil {
		return
	}
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}

func printSomething(value any) {
	intVal, ok := value.(int)

	if ok {
		fmt.Println("Integer: ", intVal)
		return
	}

	floatVal, ok := value.(float64)

	if ok {
		fmt.Println("Float: ", floatVal)
		return
	}

	stringVal, ok := value.(string)

	if ok {
		fmt.Println("String: ", stringVal)
		return
	}

	//switch value.(type) {
	//case int:
	//	fmt.Println("Integer: ", value)
	//case float64:
	//	fmt.Println("Float: ", value)
	//case string:
	//	fmt.Println(value)
	//}
}

func outputData(data outputter) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the todo failed.")
		return err
	}

	fmt.Println("Saving the todo succeeded!")
	return nil
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
