package command

import (
	"errors"
	"flag"
	"fmt"
	"github.com/mkashifaslam/golang/todo-app/action"
	"github.com/mkashifaslam/golang/todo-app/utils"
	"os"
)

func Setup() (*Command, error) {
	var err error
	osArg := os.Args

	tasksCmd := flag.NewFlagSet(string(Tasks), flag.ExitOnError)

	listCmd := flag.NewFlagSet(string(action.List), flag.ExitOnError)

	if len(osArg) < 2 || osArg[1] != string(Tasks) {
		return nil, errors.New("expected '" + string(Tasks) + "' command")
	}

	err = tasksCmd.Parse(osArg[2:])
	if err != nil {
		return nil, errors.New("invalid arguments")
	}

	args := tasksCmd.Args()

	if len(args) < 1 {
		return nil, errors.New("expected a subcommand like '" + string(action.List) + "' under '" + string(Tasks) + "'")
	}

	var cmdArg *string
	var act Act

	switch args[0] {
	case string(action.List):
		act = action.List
		err = listCmd.Parse(args)
		if err != nil {
			err = utils.ErrorHandler(err, "ListCmdFailed")
		}
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
		err = utils.NewError(fmt.Sprintf("Unknown subcommand: %s\n", args[0]))
	}

	if cmdArg == nil {
		defaultVal := ""
		cmdArg = &defaultVal
	}

	return New(Tasks, act, *cmdArg), err
}

func buildCommand(args []string, act Act) (*string, error) {
	var (
		val *string
		err error
	)

	if len(args) < 1 {
		return nil, errors.New("expected a subcommand")
	}

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

	if len(args) < 2 {
		return nil, errors.New("expected a flag argument")
	}

	err = cmd.Parse(args[1:])
	if err != nil {
		return nil, errors.New("expected a flag")
	}

	return val, err
}

func PrintHelp() {
	fmt.Println("Welcome Todos app")
	fmt.Println("----------------------")
	fmt.Printf("\nCommands:" +
		"\nhelp get help" +
		"\nexit exit app\n" +
		"\ntasks add <title>" +
		"\ntasks list" +
		"\ntasks find <taskId>" +
		"\ntasks complete <taskId>" +
		"\ntask delete <taskId>" +
		"\n\n")
	fmt.Println("----------------------")
}
