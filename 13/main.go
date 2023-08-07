package main

import "fmt"

func main() {
	a, b := 5, 6
	a, b = b, a
	fmt.Println(a, b)
}
