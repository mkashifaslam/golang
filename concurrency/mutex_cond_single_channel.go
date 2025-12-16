package main

import (
	"fmt"
	"sync"
)

func OrderWithMutex() {
	ch := make(chan int, 5)

	var wgSend sync.WaitGroup
	var wgRecv sync.WaitGroup

	// shared state for ordered printing
	expected := 1
	mu := &sync.Mutex{}
	cond := sync.NewCond(mu)

	// --- 5 senders ---
	wgSend.Add(5)
	for i := 1; i <= 5; i++ {
		go func(val int) {
			defer wgSend.Done()
			ch <- val
		}(i)
	}

	// Close channel once sends complete
	go func() {
		wgSend.Wait()
		close(ch)
	}()

	// --- 5 receivers ---
	wgRecv.Add(5)
	for i := 1; i <= 5; i++ {
		go func(id int) {
			fmt.Printf("Receiver %d started\n", id) // debug print
			defer wgRecv.Done()
			for v := range ch {
				fmt.Printf("Receiver %d received: %d\n", id, v) // debug print
				mu.Lock()
				for v != expected { // wait until correct number
					fmt.Printf("Receiver %d waiting for: %d\n", id, expected) // debug print
					cond.Wait()
				}

				fmt.Printf("Receiver %d got: %d\n", id, v)
				expected++
				cond.Broadcast() // wake up others
				mu.Unlock()
				break
			}
			fmt.Printf("Receiver %d done\n", id) // debug print
		}(i)
	}

	wgRecv.Wait()
}
