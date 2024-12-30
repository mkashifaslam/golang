package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserCommand() string {
	var command string
	GetUserInput(&command)
	return command
}

func GetUserInput(input *string) {
	err := GetUserInputCmd(input)
	if err != nil {
		fmt.Println("UserInput: user input was incorrect")
		fmt.Println("Error detail", err)
	}
}

func GetUserInputCmd(input *string) error {
	_, err := fmt.Scan(input)
	if err != nil {
		return err
	}

	return nil
}

func GetTitle(prompt string) string {
	title := GetUserInputReader(prompt)
	if title == "" {
		fmt.Println("TitleInput: user input was incorrect")
	}
	return title
}

func GetUserInputReader(prompt string) string {
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
