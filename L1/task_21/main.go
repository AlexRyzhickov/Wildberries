package main

import "fmt"

type ExternalInterface interface {
	targetFunc()
}

type InternalStruct struct {
}

func (i InternalStruct) someRequiredFunc() {
	fmt.Println("Hi from someRequiredFunc !")
}

type Adapter struct {
	i InternalStruct
}

func (a Adapter) targetFunc() {
	a.i.someRequiredFunc()
}

func main() {
	var e ExternalInterface = Adapter{i: InternalStruct{}}
	e.targetFunc()
}
