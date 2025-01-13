package main

import (
	"context"
	"fmt"
	"time"
)

func cleanup() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("cleanup and exit")
				return
			default:
				fmt.Println("1 second")
				time.Sleep(time.Second)
			}
		}
	}(ctx)

	time.Sleep(5 * time.Second)
}
