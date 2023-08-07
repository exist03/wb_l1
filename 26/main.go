package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(checkString("abcdA"))
	fmt.Println(checkString("abcd"))
}
func checkString(str string) bool {
	sl := []rune(strings.ToLower(str))
	myMap := make(map[rune]struct{})
	for _, v := range sl {
		if _, ok := myMap[v]; !ok {
			myMap[v] = struct{}{}
		} else {
			return false
		}
	}
	return true
}
