package app

import (
	"fmt"
	a "github.com/mkashifaslam/golang/todo-app/action"
)

func Run(action a.Act, args string) {
	a.Run(action, args)

	if action == a.List {
		fmt.Println("--------------------------------------------")
	}

	fmt.Println("Done")
}
