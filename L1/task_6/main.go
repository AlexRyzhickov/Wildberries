package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	quitChannel := make(chan struct{})
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func(ctx context.Context, quit <-chan struct{}) {
		for {
			time.Sleep(time.Millisecond * 100)
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine stopped work by context.WithTimeout")
				return
			case <-quit:
				fmt.Println("Goroutine stopped work by quit channel")
				return
			default:
				fmt.Println("Msg from goroutine")
			}
		}
	}(ctx, quitChannel)

	//quitChannel <- struct{}{}
	time.Sleep(3 * time.Second)
}
