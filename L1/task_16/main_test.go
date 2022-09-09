package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	assert.Equal(t, binarySearch(items, 63), 6, "they should be equal")
	assert.Equal(t, binarySearch(items, 1), 0, "they should be equal")
	assert.Equal(t, binarySearch(items, 100), len(items)-1, "they should be equal")
	assert.Equal(t, binarySearch(items, 11), -1, "they should be equal")
}
