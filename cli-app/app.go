package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	tasksCmd := flag.NewFlagSet("tasks", flag.ExitOnError)
	token := tasksCmd.String("t", "", "token")

	// os.Args[0] = main program
	// os.Args[1] = command
	// os.Args[2:] = arguments

	fmt.Printf("Main Program: %s \nCommand: %s \nArgs: %v \n", os.Args[0], os.Args[1], os.Args[2:])
	args := os.Args[2:]

	err := tasksCmd.Parse(args)

	if err != nil {
		fmt.Printf("Error parsing flags: %v\n", err)
		return
	}

	//flag.Parse()

	if *token == "" {
		err := fmt.Errorf("token is required")
		fmt.Printf("Err: %s\n", err.Error())
		return
	}

	fmt.Println("token:", *token)
}
