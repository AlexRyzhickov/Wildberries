package main

import (
	"fmt"
	"sync"
)

func square(wg *sync.WaitGroup, i int, limit chan<- int) {
	wg.Done()
	limit <- i * i
}

func main() {
	sum := 0
	results := make(chan int, 3)
	nums := []int{2, 4, 6, 8, 10}

	wg := sync.WaitGroup{}
	wg.Add(len(nums))

	for _, task := range nums {
		go square(&wg, task, results)
	}

	wg.Wait()

	for i := 0; i < len(nums); i++ {
		sum += <-results
	}
	close(results)

	fmt.Println("Result: ", sum)
}
