package calculator

import (
	"Wildberries/L2/patterns/structural/facade/notation"
	"Wildberries/L2/patterns/structural/facade/stack"
	"strconv"
)

var operations = map[string]func(a float64, b float64) float64{
	"+": func(a float64, b float64) float64 { return a + b },
	"-": func(a float64, b float64) float64 { return a - b },
	"*": func(a float64, b float64) float64 { return a * b },
	"/": func(a float64, b float64) float64 { return a / b },
}

type calculator struct{}

type Calc interface {
	Calculate(expression string) float64
}

func NewCalculator() Calc {
	return &calculator{}
}

func (c calculator) Calculate(expression string) float64 {
	postfix := notation.FromInfixToPostfix(expression)

	var stack stack.FloatStack

	for _, v := range postfix {
		if num, err := strconv.Atoi(v); err == nil {
			stack.Push(float64(num))
		} else if num, err := strconv.ParseFloat(v, 8); err == nil {
			stack.Push(num)
		} else {
			if v != " " {
				a, e1 := stack.Pop()
				b, e2 := stack.Pop()

				if e1 == true && e2 == true {
					val := operations[v](b, a)
					stack.Push(val)
				}
			}
		}
	}

	pop, _ := stack.Pop()

	return pop
}
