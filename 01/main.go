package main

import "fmt"

type Human struct {
	filed1 int
	field2 string
}

func (h Human) f() {
	fmt.Println(h.field2)
}

type Action struct {
	Human
}

func main() {
	a := Action{Human{
		filed1: 0,
		field2: "smth",
	}}
	a.f()
}
