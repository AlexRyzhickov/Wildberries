package main

import (
	"fmt"
)

func intersectionOfSets[T comparable](set1 []T, set2 []T) []T {
	m := make(map[T]struct{})
	intersection := make([]T, 0)
	for _, v := range set1 {
		m[v] = struct{}{}
	}
	for _, v := range set2 {
		_, ok := m[v]
		if ok {
			intersection = append(intersection, v)
		}
	}
	return intersection
}

func main() {
	nums1 := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	nums2 := []float64{-25.8, 417, 122, 1, 19, 13.0, 19.0, 15.4, 24.5, -21.7, 32.5, 100}
	fmt.Println(intersectionOfSets(nums1, nums2))

	strs1 := []string{"hello", "world", "!"}
	strs2 := []string{"world", "goodbye", "!!!"}
	fmt.Println(intersectionOfSets(strs1, strs2))
}
