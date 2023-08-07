package main

import "fmt"

func main() {
	a := 0.
	detectType(a)
}
func detectType(a interface{}) {
	switch v := a.(type) {
	default:
		fmt.Printf("%T", v)
	}
}
