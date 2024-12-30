package main

import (
	"github.com/mkashifaslam/golang/todo-app/command"
	"github.com/mkashifaslam/golang/todo-app/input"
)

func main() {
	for {
		userCmd := input.GetUserCommand()
		command.Handler(userCmd)
	}
}
