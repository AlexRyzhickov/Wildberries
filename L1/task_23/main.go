package main

import (
	"fmt"
)

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	slice := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(remove(slice, 2))
}
