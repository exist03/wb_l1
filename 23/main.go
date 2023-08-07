package main

import "fmt"

func trim[T any](sl []T, ind int) []T {
	return append(sl[:ind], sl[ind+1:]...)
}

func main() {
	sl := []int{0, 1, 2, 3, 4}
	fmt.Println(sl)
	sl = trim(sl, 3)
	fmt.Println(sl)
}
