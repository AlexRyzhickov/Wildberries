package main

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	ints := []int{-25, -27, 32, -41, 15, 13, 19, 15, 24, -21.0}
	copyInts := make([]int, len(ints))
	copy(copyInts, ints)
	copyInts = quicksort(copyInts)
	sort.Ints(ints)
	for i := 0; i < len(ints); i++ {
		assert.Equal(t, copyInts[i], ints[i], "they should be equal")
	}

	floats := []float64{-25, -27, 32, -41, 15, 13, 19, 15, 24, -21.0}
	copyFloats := make([]float64, len(floats))
	copy(copyFloats, floats)
	copyFloats = quicksort(copyFloats)
	sort.Float64s(floats)
	for i := 0; i < len(floats); i++ {
		assert.Equal(t, copyFloats[i], floats[i], "they should be equal")
	}
}
