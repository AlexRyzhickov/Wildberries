package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ch := make(chan int, 10)

	fmt.Println("Enter the number of workers")

	var numberWorkers int
	fmt.Scanf("%d\n", &numberWorkers)

	for i := 0; i < numberWorkers; i++ {
		go func(c <-chan int) {
			for number := range c {
				fmt.Println(number)
			}
		}(ch)
	}

	isAlways := true

	go func() {
		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
		defer signal.Stop(shutdown)
		<-shutdown
		isAlways = false
	}()

	for isAlways {
		time.Sleep(100 * time.Millisecond)
		rand.Seed(time.Now().UnixNano())
		number := rand.Intn(100)
		ch <- number
	}
	close(ch)
}
