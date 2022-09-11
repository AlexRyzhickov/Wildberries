package main

import (
	"fmt"
	"sync"
)

func square(wg *sync.WaitGroup, i int) {
	fmt.Println(i * i)
	wg.Done()
}

func main() {
	nums := []int{2, 4, 6, 8, 10}

	wg := sync.WaitGroup{}
	wg.Add(len(nums))

	for _, task := range nums {
		go square(&wg, task)
	}

	wg.Wait()
}
