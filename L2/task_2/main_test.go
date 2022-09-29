package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnpackString(t *testing.T) {
	unpackString("")
	assert.Equal(t, unpackString("я4яcяd5e"), "яяяяяcяddddde", "they should be equal")
	assert.Equal(t, unpackString("a1b5c2d5e11"), "abbbbbccdddddeeeeeeeeeee", "they should be equal")
	assert.Equal(t, unpackString("a1b5c2d5e"), "abbbbbccddddde", "they should be equal")
	assert.Equal(t, unpackString("a2d3b4"), "aadddbbbb", "they should be equal")
	assert.Equal(t, unpackString("a4"), "aaaa", "they should be equal")
	assert.Equal(t, unpackString("a"), "a", "they should be equal")
	assert.Equal(t, unpackString("4a"), "a", "they should be equal")
	assert.Equal(t, unpackString("45"), "", "they should be equal")
	assert.Equal(t, unpackString("4"), "", "they should be equal")
	assert.Equal(t, unpackString(""), "", "they should be equal")
}
