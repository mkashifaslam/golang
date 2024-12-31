package app

import (
	"fmt"
	a "github.com/mkashifaslam/golang/todo-app/action"
)

func Run(action a.Act, args string) {
	a.Run(action, args)
	fmt.Println("Done")
}
