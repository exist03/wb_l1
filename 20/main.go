package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog   sun"
	sl := strings.Fields(str)
	for i := len(sl) - 1; i >= 0; i-- {
		fmt.Print(sl[i])
		if i != 0 {
			fmt.Print(" ")
		}
	}
}
