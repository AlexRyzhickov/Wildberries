package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i interface{}

	i = 1
	//i = true
	//i = make(chan int)

	//1. Method to determine the type
	t := reflect.TypeOf(i)
	fmt.Println(t)

	//2. Method to determine the type
	switch i.(type) {
	case int:
		fmt.Println("int")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan int")
	default:
		fmt.Println("unknown type")
	}

	//3. Method to determine the type
	fmt.Println(fmt.Sprintf("%T", i))
}
