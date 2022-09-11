package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func binarySearch[T Number](a []T, key T) int {
	l := -1
	r := len(a)
	for l < r-1 {
		m := (l + r) / 2
		if a[m] < key {
			l = m
		} else {
			r = m
		}
	}
	if a[r] != key {
		return -1
	}
	return r
}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	if v := binarySearch(items, 20); v != -1 {
		fmt.Println("Element found, position", v)
	} else {
		fmt.Println("Element not found")
	}
}
