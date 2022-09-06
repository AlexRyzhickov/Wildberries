package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) sayYourName() {
	fmt.Println("My name", h.Name)
}

func (h Human) sayYourAge() {
	fmt.Println("My age", h.Age)
}

type Action struct {
	Human
}

func main() {
	a := Action{}
	a.Name = "Alex"
	a.Age = 22
	a.sayYourAge()
	a.sayYourAge()
}
