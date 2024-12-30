package utils

import (
	"errors"
	"fmt"
)

func ErrorHandler(err error, msg string) error {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return GetErrorMessage(err, msg)
	}

	return nil
}

func GetErrorMessage(err error, msg string) (error error) {
	if msg != "" {
		error = errors.New(msg)
	} else if err != nil {
		error = errors.New("Error: " + err.Error())
	} else {
		error = errors.New("unknown error")
	}

	return
}

func PrintError(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err.Error())
	}
}
