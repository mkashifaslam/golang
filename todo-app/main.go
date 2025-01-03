package main

import (
	"github.com/mkashifaslam/golang/todo-app/command"
)

func main() {
	command.PrintHelp()

	for {
		command.Run()
	}
}
