package utils

import (
	"errors"
	"fmt"
	"math/rand"
)

func NewError(msg string) error {
	return errors.New(msg)
}

func ErrorHandler(err error, msg string) error {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return getErrorMessage(err, msg)
	}

	return nil
}

func getErrorMessage(err error, msg string) (error error) {
	if msg != "" {
		error = NewError(msg)
	} else if err != nil {
		error = NewError("Error: " + err.Error())
	} else {
		error = NewError("unknown error")
	}

	return
}

func PrintError(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err.Error())
	}
}

func GenerateRandomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}
