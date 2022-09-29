package main

import (
	"fmt"
	"strings"
	"unicode"
)

func scanNumber(runes []rune, pos int, len int) (int, int) {
	count := int(runes[pos] - '0')
	step := pos
	for i := pos + 1; i < len; i++ {
		if unicode.IsDigit(runes[i]) {
			count = count*10 + int(runes[i]-'0')
			step = i
		} else {
			break
		}
	}
	return count, step - pos
}

func unpackString(str string) string {
	var sb strings.Builder
	runes := []rune(str)
	last := len(runes)
	for i := 0; i < last-1; i++ {
		if unicode.IsLetter(runes[i]) {
			if unicode.IsDigit(runes[i+1]) {
				count, step := scanNumber(runes, i+1, last)
				for j := 0; j < count; j++ {
					sb.WriteString(string(runes[i]))
				}
				i = i + step
			} else {
				sb.WriteString(string(runes[i]))
			}
		}
	}
	if last > 0 && unicode.IsLetter(runes[last-1]) {
		sb.WriteString(string(runes[last-1]))
	}
	return sb.String()
}

func main() {
	fmt.Println(unpackString("a1b5c2d5e11"))
	fmt.Println(unpackString("a1b5c2d5e"))
}
