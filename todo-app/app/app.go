package app

import (
	"fmt"
	"github.com/mkashifaslam/golang/todo-app/action"
)

func Run() {
	// take user action input
	actionInput := action.GetActionInput()
	// execute user action
	action.Run(actionInput)
	fmt.Println("Done")
}
