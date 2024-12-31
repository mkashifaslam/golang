package main

import (
	"fmt"
	cmd "github.com/mkashifaslam/golang/todo-app/command"
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
		userCmd, err := cmd.GetInput()
		utils.PrintError(err, "InputError:")
		inputCmd, err := cmd.Parse(userCmd)
		utils.PrintError(err, "ParsingError:")
		if inputCmd != nil {
			cmd.Handler(inputCmd.App, inputCmd.Action, inputCmd.Args)
		}
	}
}
