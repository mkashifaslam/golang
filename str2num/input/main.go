package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetStringInput(prompt string) string {
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
