package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	go func() {
		for number := range ch {
			fmt.Println("x =", number, "x * x =", number*number)
		}
	}()

	for n := range nums {
		ch <- n
	}
	close(ch)
}
