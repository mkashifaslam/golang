package main

import (
	"github.com/mkashifaslam/golang/todo-app/command"
	"github.com/mkashifaslam/golang/todo-app/input"
	"github.com/mkashifaslam/golang/todo-app/utils"
)

func main() {
	for {
		userCmd, err := input.GetUserInput()
		utils.PrintError(err, "Error reading command")
		command.Handler(userCmd)
	}
}
