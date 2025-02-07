package main

import (
	"fmt"
	"strconv"
)

func main() {
	err := Str2Num()
	if err != nil {
		fmt.Println("Error happened:", err)
	}
}

func Str2Num() (err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
			}
			return
		}
	}()

	var str string
	fmt.Printf("Enter a number: ")
	fmt.Scan(&str)
	result := str2Num(str)
	fmt.Printf("Result: %d\n", result)

	return
}

func str2Num(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}
