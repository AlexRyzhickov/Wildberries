package main

import (
	"fmt"
	"strings"
)

func reverseWordsInString(str string) string {
	strs := strings.Split(str, " ")
	for i := 0; i < len(strs)/2; i++ {
		strs[i], strs[len(strs)-i-1] = strs[len(strs)-i-1], strs[i]
	}
	return strings.Join(strs, " ")
}

func main() {
	fmt.Println(reverseWordsInString("snow dog sun"))
}
