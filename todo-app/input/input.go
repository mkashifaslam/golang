package input

import (
	"bufio"
	"fmt"
	"github.com/mkashifaslam/golang/todo-app/utils"
	"os"
	"strings"
)

func GetUserInput() (string, error) {
	var input string
	err := getCmdInput(&input)
	formattedError := utils.ErrorHandler(err, "please enter a valid user input")

	return input, formattedError
}

func getCmdInput(input *string) error {
	_, err := fmt.Scan(input)
	if err != nil {
		return err
	}

	return nil
}

func GetUserStringInput(prompt string) string {
	input := getStringInput(prompt)
	if input == "" {
		fmt.Println("UserInput: user input was incorrect")
	}
	return input
}

func getStringInput(prompt string) string {
	fmt.Printf("%s ", prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
