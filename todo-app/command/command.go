package command

import (
	"fmt"
	"github.com/mkashifaslam/golang/todo-app/action"
	"github.com/mkashifaslam/golang/todo-app/app"
	"github.com/mkashifaslam/golang/todo-app/input"
	"github.com/mkashifaslam/golang/todo-app/utils"
	"os"
	"strings"
)

type Cmd string

type Act = action.Act

const (
	Tasks Cmd = "tasks"
	Help  Cmd = "help"
	Exit  Cmd = "exit"
)

type Command struct {
	App    Cmd
	Action Act
	Args   string
}

func (cmd *Command) Print() {
	fmt.Printf("%v\n %v\n %s\n", cmd.App, cmd.Action, cmd.Args)
}

func New(app Cmd, action Act, args string) *Command {
	return &Command{App: app, Action: action, Args: args}
}

func Handler(cmd Cmd, action Act, args string) {
	switch cmd {
	case Tasks:
		app.Run(action, args)
	case Help:
		fmt.Println("Help: tasks <command> [<args>]")
	case Exit:
		os.Exit(0)
	default:
		fmt.Println("Help: tasks <command> [<args>]")
	}
}

func GetInput() (string, error) {
	cmd := input.GetUserStringInput("todo-app$:")
	if cmd == "" {
		return "", utils.ErrorHandler(nil, "Command not valid")
	}
	return cmd, nil
}

func Parse(cmd string) (*Command, error) {
	var (
		cmdApp Cmd
		cmdAct Act
		cmdArg string
	)
	cmdSlice := strings.Split(cmd, " ")

	if len(cmdSlice) < 1 {
		return nil, utils.NewError("Command not found")
	}

	cmdApp = Cmd(cmdSlice[0])

	if cmdApp != Tasks {
		return New(cmdApp, cmdAct, cmdArg), nil
	}

	if len(cmdSlice) < 2 {
		return nil, utils.NewError("Command not valid")
	}

	cmdApp, cmdAct = Cmd(cmdSlice[0]), Act(cmdSlice[1])

	if len(cmdSlice) > 2 {
		cmdArg = strings.Join(cmdSlice[2:], " ")
	}

	if isArgsRequired(cmdAct) && cmdArg == "" {
		return nil, utils.NewError("Command args not valid")
	}

	return New(cmdApp, cmdAct, cmdArg), nil
}

func isArgsRequired(cmdAction Act) bool {
	return action.IsValid(cmdAction) && cmdAction != action.List
}
