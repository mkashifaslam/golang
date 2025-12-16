package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Channel() {

	var sWg sync.WaitGroup
	var rWg sync.WaitGroup

	num := 5

	ch := make(chan int, num)

	sendChanValue(ch, num, &sWg)

	go func() {
		sWg.Wait()
		close(ch)
	}()

	//receiveChanWithSelect(ch)

	//receiveChanWithFor(ch, num)

	//receiveChanWithSelect(ch)

	receiveChanWithWait(ch, num, &rWg)

	rWg.Wait()

	fmt.Println("All goroutines done")

	//time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	//close(ch)

}

func sendChanValue(ch chan int, num int, wg *sync.WaitGroup) {
	wg.Add(num)
	for i := 1; i <= num; i++ {
		go func(i int) {
			defer wg.Done()
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
	var random = rand.New(rand.NewSource(time.Now().UnixNano())).Intn(10)
	var ttl = time.Duration(random) * time.Millisecond
	fmt.Printf("timeout set to %ds, random number is %d \n", ttl/1000, random)
	timeout := time.After(ttl)

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
	wg.Add(num)
	for i := 1; i <= num; i++ {
		go func() {
			defer wg.Done()
			fmt.Println("goroutine output", <-ch)
		}()
	}
}
