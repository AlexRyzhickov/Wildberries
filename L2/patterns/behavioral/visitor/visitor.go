package main

import (
	"fmt"
	"math"
)

type GeometricShapes interface {
	accept(Visitor)
}

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	w float64
	h float64
}

func (s *Square) accept(v Visitor) {
	v.visitForSquare(s)
}

func (c *Circle) accept(v Visitor) {
	v.visitForCircle(c)
}

func (r *Rectangle) accept(v Visitor) {
	v.visitForRectangle(r)
}

type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForRectangle(*Rectangle)
}

type AreaCalculator struct {
	area float64
}

func (a *AreaCalculator) visitForSquare(s *Square) {
	a.area = s.side * s.side
	fmt.Println("Square area:", a.area)
}

func (a *AreaCalculator) visitForCircle(s *Circle) {
	a.area = math.Pi * s.radius * s.radius
	fmt.Println("Circle area:", a.area)
}

func (a *AreaCalculator) visitForRectangle(s *Rectangle) {
	a.area = s.w * s.h
	fmt.Println("Rectangle area:", a.area)
}

func main() {
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{w: 2, h: 3}

	areaCalculator := &AreaCalculator{}
	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)
}

/*
Посетитель — это поведенческий паттерн, который позволяет добавить новую операцию для целой иерархии классов, не изменяя код этих классов.
+ Позволяет добавить новый метод для класса, без изменения самого класса.
+ Используется для добавления нового поведение только для некоторых классов из существующих.
*/
