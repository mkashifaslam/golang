package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	go greet("Greeting # 1", done)
	go greet("Greeting # 2", done)
	go slowGreet("Greeting # 4 - slow", done)
	go greet("Greeting # 3", done)

	for range done {
	}
}

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(time.Second * 3)
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan)
}
