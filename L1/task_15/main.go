package main

//Проблемой указаного ниже кода является то,
//что создаётся строка большого размера 1 << 10 и при
//этом берутся только первые 100 символов, при этом все элементы
//строки от [100 : 1023] остаются в памяти после нарезки и не используются.

/*
var justString string
func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}*/

//Правильным подходом является использование функции Clone (аналог copy)

/*
import "strings"

var justString string
func someFunc() {
	v := createHugeString(1 << 10)
	justString = strings.Clone(v[:100])
}

func main() {
	someFunc()
}*/
