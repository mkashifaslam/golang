package command

import (
	"fmt"
	"github.com/mkashifaslam/golang/todo-app/action"
	"github.com/mkashifaslam/golang/todo-app/app"
)

type Cmd string

type Act = action.Act

type Command struct {
	Action Act
	Args   string
}

func (cmd *Command) Print() {
	fmt.Printf("%v\n %s\n", cmd.Action, cmd.Args)
}

func New(action Act, args string) *Command {
	return &Command{Action: action, Args: args}
}

func Run() {
	PrintHelp()
	cmd, err := Setup()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Handler(cmd.Action, cmd.Args)
}

func Handler(action Act, args string) {
	app.Run(action, args)
}
