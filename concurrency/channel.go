package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Channel() {

	var wg sync.WaitGroup

	num := 5

	ch := make(chan int)

	sendChanValue(ch, num, &wg)

	//receiveChanWithSelect(ch)

	//receiveChanWithFor(ch, num)

	//receiveChanWithSelect(ch)

	receiveChanWithWait(ch, num, &wg)

	wg.Wait()

	fmt.Println("All goroutines done")

	//time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	//close(ch)

}

func sendChanValue(ch chan int, num int, wg *sync.WaitGroup) {
	for i := 1; i <= num; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Printf("goroutine input %d\n", i)
			ch <- i
		}(i)
	}

}

func receiveChanWithFor(ch chan int, num int) {
	for i := 1; i <= num; i++ {
		fmt.Println("output", <-ch)
	}
}

func receiveChanWithSelect(ch chan int) {
	timeout := time.After(time.Duration(rand.Intn(100)) * time.Millisecond)

	for {
		select {
		case i := <-ch:
			fmt.Printf("goroutine output %d\n", i)
		case <-timeout:
			fmt.Println("timeout")
			close(ch)
			return
		}
	}
}

func receiveChanWithWait(ch chan int, num int, wg *sync.WaitGroup) {
	for i := 1; i <= num; i++ {
		go func() {
			defer wg.Done()
			fmt.Println("goroutine output", <-ch)
		}()
	}

}
