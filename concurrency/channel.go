package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Channel() {
	num := 5

	ch := make(chan int)

	sendChanValue(ch, num)

	//receiveChanWithFor(ch, num)

	receiveChanWithSelect(ch)

}

func sendChanValue(ch chan int, num int) {
	for i := 1; i <= num; i++ {
		go func(i int) {
			fmt.Printf("goroutine %d\n", i)
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
	timeout := time.After(time.Duration(rand.Intn(5)) * time.Millisecond)

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
