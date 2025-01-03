package command

import "fmt"

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
