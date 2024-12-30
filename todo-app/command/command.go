package command

import (
	"fmt"
	"github.com/mkashifaslam/golang/todo-app/app"
	"os"
)

const (
	MainCmd = "tasks"
	HelpCmd = "help"
	ExitCmd = "exit"
)

func Handler(cmd string) {
	switch cmd {
	case MainCmd:
		app.Run()
	case HelpCmd:
		fmt.Println("Help: tasks <command> [<args>]")
	case ExitCmd:
		os.Exit(1)
	default:
		fmt.Println("Help: tasks <command> [<args>]")
	}
}
