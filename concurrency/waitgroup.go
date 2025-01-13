package main

import (
	"fmt"
	"sync"
)

func waitGroup() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Worker %d doing work\n", id)
		}(i)
	}

	wg.Wait()

	fmt.Println("All work finished")
}
