package command

import (
	"errors"
	"flag"
	"fmt"
	"github.com/mkashifaslam/golang/todo-app/action"
	"github.com/mkashifaslam/golang/todo-app/utils"
	"os"
)

func Setup() (cmd *Command, err error) {
	osArgs := os.Args

	if len(osArgs) < 2 {
		err = errors.New("no command specified")
		return
	}

	args := osArgs[2:]

	var (
		cmdArg *string
		act    Act
	)

	switch osArgs[1] {
	case string(action.List):
		act = action.List
		cmdArg = nil
	case string(action.Add):
		act = action.Add
		cmdArg, err = buildCommand(args, action.Add)
		if err != nil {
			err = utils.ErrorHandler(err, "AddCmdFailed")
		}
	case string(action.Find):
		act = action.Find
		cmdArg, err = buildCommand(args, action.Find)
		if err != nil {
			err = utils.ErrorHandler(err, "FindCmdFailed")
		}
	case string(action.Complete):
		act = action.Complete
		cmdArg, err = buildCommand(args, action.Complete)
		if err != nil {
			err = utils.ErrorHandler(err, "CompleteCmdFailed")
		}
	case string(action.Delete):
		act = action.Delete
		cmdArg, err = buildCommand(args, action.Delete)
		if err != nil {
			err = utils.ErrorHandler(err, "DeleteCmdFailed")
		}
	default:
		err = utils.ErrorHandler(err, "CmdFailed")
	}

	if cmdArg == nil {
		defaultVal := ""
		cmdArg = &defaultVal
	}

	cmd = New(act, *cmdArg)

	return
}

func buildCommand(args []string, act Act) (*string, error) {
	var (
		val *string
		err error
	)

	cmd := flag.NewFlagSet(string(act), flag.ExitOnError)

	switch act {
	case action.Add:
		val = cmd.String("title", "", "Title of the task")
	case action.Find:
		val = cmd.String("id", "", "ID of the task")
	case action.Complete:
		val = cmd.String("id", "", "ID of the task")
	case action.Delete:
		val = cmd.String("id", "", "ID of the task")
	}

	if len(args) < 1 {
		return nil, errors.New("expected a flag argument")
	}

	if len(args) < 2 {
		return nil, errors.New("invalid flag argument")
	}

	err = cmd.Parse(args)
	if err != nil {
		return nil, errors.New("flag parsing error")
	}

	return val, err
}

func PrintHelp() {
	fmt.Println("Welcome Todos app")
	fmt.Println("--------------------------------------------")
	fmt.Printf("\nCommands:" +
		"\nlist" +
		"\nfind <taskId>" +
		"\ncomplete <taskId>" +
		"\ndelete <taskId>" +
		"\n\n")
	fmt.Println("--------------------------------------------")
}
