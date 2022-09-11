package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	intItems := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	assert.Equal(t, binarySearch(intItems, 63), 6, "they should be equal")
	assert.Equal(t, binarySearch(intItems, 1), 0, "they should be equal")
	assert.Equal(t, binarySearch(intItems, 100), len(intItems)-1, "they should be equal")
	assert.Equal(t, binarySearch(intItems, 11), -1, "they should be equal")

	floatItems := []float32{1.4, 2.5, 9.11, 20.005, 31.14, 45.16, 63.5, 70.45, 100.5}
	assert.Equal(t, binarySearch(floatItems, 63.5), 6, "they should be equal")
	assert.Equal(t, binarySearch(floatItems, 1.4), 0, "they should be equal")
	assert.Equal(t, binarySearch(floatItems, 100.5), len(intItems)-1, "they should be equal")
	assert.Equal(t, binarySearch(floatItems, 11.17), -1, "they should be equal")
}
