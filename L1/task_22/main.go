package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1 << 21)
	b := big.NewInt(1 << 22)
	res := new(big.Int)

	fmt.Println("a =", a, "b =", b)
	fmt.Println("Sum: a + b =", res.Add(a, b))
	fmt.Println("Sub: a - b =", res.Sub(a, b))
	fmt.Println("Div: b / a =", res.Div(b, a))
	fmt.Println("Mul: a * b =", res.Mul(a, b))
}
