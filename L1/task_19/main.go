package main

import "fmt"

func reverseString(str string) string {
	runes := []rune(str)
	var res []rune
	for i := len(runes) - 1; i >= 0; i-- {
		res = append(res, runes[i])
	}
	return string(res)
}

func main() {
	fmt.Println(reverseString("Hello, 愚公山を移す"))
	fmt.Println(reverseString("Довод"))
	fmt.Println(reverseString("Tenet"))
}
