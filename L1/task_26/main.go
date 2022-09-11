package main

import (
	"fmt"
	"strings"
)

func characterUniquenessCheck(str string) bool {
	m := make(map[rune]int)
	for _, v := range strings.ToLower(str) {
		m[v] = m[v] + 1
	}
	for _, v := range m {
		if v != 1 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(characterUniquenessCheck("abcd"))
	fmt.Println(characterUniquenessCheck("abCdefAaf"))
	fmt.Println(characterUniquenessCheck("aabcd"))
}
