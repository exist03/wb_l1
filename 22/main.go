package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)
	res := new(big.Int)
	for {
		var str string
		fmt.Print("Enter first number: ")
		fmt.Scan(&str)
		if _, ok := a.SetString(str, 10); ok {
			break
		}
	}
	for {
		var str string
		fmt.Print("Enter second number: ")
		fmt.Scan(&str)
		if _, ok := b.SetString(str, 10); ok {
			break
		}
	}
	var operation string
	fmt.Print("Enter the operation: ")
	fmt.Scan(&operation)
	switch operation {
	case "/":
		res.Div(a, b)
	case "*":
		res.Mul(a, b)
	case "+":
		res.Add(a, b)
	case "-":
		res.Sub(a, b)
	}
	fmt.Println(res)
}
