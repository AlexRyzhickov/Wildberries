package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"math/rand"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func quicksort[T Number](slice []T) []T {
	if len(slice) < 2 {
		return slice
	}
	left, right := 0, len(slice)-1
	pivot := rand.Int() % len(slice)
	slice[pivot], slice[right] = slice[right], slice[pivot]
	for i, _ := range slice {
		if slice[i] < slice[right] {
			slice[left], slice[i] = slice[i], slice[left]
			left++
		}
	}
	slice[left], slice[right] = slice[right], slice[left]
	quicksort(slice[:left])
	quicksort(slice[left+1:])
	return slice
}

func main() {
	slice := []int{-25, -27, 32, -41, 15, 13, 19, 15, 24, -21.0}
	fmt.Println(quicksort(slice))
}
