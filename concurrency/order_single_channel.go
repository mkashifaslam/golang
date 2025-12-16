package main

import (
	"fmt"
	"sync"
)

func OrderWithSingleChannel() {
	ch := make(chan int, 5) // buffer ensures sends aren't competing randomly

	var wgSend sync.WaitGroup
	var wgRecv sync.WaitGroup

	// 5 Sender goroutines
	wgSend.Add(5)
	for i := 1; i <= 5; i++ {
		go func(val int) {
			defer wgSend.Done()
			ch <- val // send numbers concurrently
		}(i)
	}

	// Close channel when all sends are done
	go func() {
		wgSend.Wait()
		close(ch)
	}()

	// 5 receiver goroutines but sequentialized using range
	wgRecv.Add(5)
	for i := 1; i <= 5; i++ {
		go func(id int) {
			defer wgRecv.Done()

			// Read values in deterministic order
			v := <-ch // receives next in channel FIFO
			fmt.Printf("Receiver %d got: %d\n", id, v)
		}(i)
	}

	wgRecv.Wait()
}
