package main

import (
	"fmt"
)

func main() {
	fmt.Println("Program to convert strings to numbers")

	// store map of string keyword to number
	// one - nine => 1 - 9
	// hundred => 100 -> 1 * 100
	// thousand => 1,000 -> 100 * 10
	// million => 1,000,000 -> 1,000* 1,000
	// billion => 1,000,000,000 -> 1,000,000 * 1,000
	// trillion => 1,000,000,000,000 -> 1,000,000,000 * 1,000
	// quadrillion => 1,000,000,000,000,000

	var str = GetStringInput("Enter a string: ")
	if str == "" {
		fmt.Println("Input string is empty")
		return
	}

	fmt.Printf("Numer is %d\n", Str2Num(str))

}
