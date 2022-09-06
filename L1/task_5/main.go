package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	timeCh := time.After(1 * time.Second)
	rand.Seed(time.Now().UnixNano())

	go func() {
		for number := range ch {
			fmt.Println("Received", number)
		}
	}()

Loop:

	for {
		select {
		case <-timeCh:
			fmt.Println("Timed out")
			break Loop
		default:
			number := rand.Intn(100)
			ch <- number
		}
	}

	close(ch)
}
