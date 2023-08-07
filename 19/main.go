package main

import "fmt"

func main() {
	str := "qwerty"
	n := reverse(str)
	fmt.Println(n)
}

func reverse(str string) string {
	sl := []rune(str)
	r := len(sl) - 1
	for l := 0; l < r; l++ {
		sl[l], sl[r] = sl[r], sl[l]
		r--
	}
	return string(sl)
}
