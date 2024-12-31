package main

import (
	"fmt"
	"github.com/mkashifaslam/golang/todo-app/command"
	"github.com/mkashifaslam/golang/todo-app/input"
	"github.com/mkashifaslam/golang/todo-app/utils"
)

func main() {
	fmt.Println("Welcome Todos app")
	fmt.Println("----------------------")
	fmt.Printf("\nCommands:" +
		"\ntasks add <title>" +
		"\ntasks list" +
		"\ntasks find <taskId>" +
		"\ntasks complete <taskId>" +
		"\ntask delete <taskId>" +
		"\n\n")
	fmt.Println("----------------------")
	for {
		userCmd, err := input.GetUserInput()
		utils.PrintError(err, "Error reading command")
		command.Handler(userCmd)
	}
}
