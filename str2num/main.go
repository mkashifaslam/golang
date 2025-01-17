package main

import (
	"fmt"
	"github.com/mkashifaslam/golang/str2num/input"
	"github.com/mkashifaslam/golang/str2num/lib"
	"github.com/mkashifaslam/golang/str2num/output"
)

func main() {
	fmt.Println("Program to convert strings to numbers")
	var str = input.GetStringInput("Enter a string: ")
	if str == "" {
		fmt.Println("Input string is empty")
		return
	}

	number := lib.Str2Num(str)
	result := output.FormatCommas(number)

	fmt.Println("Number is", result)

}
