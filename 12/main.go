package main

import "fmt"

func main() {
	sl := []string{"cat", "cat", "dog", "cat", "tree"}
	myMap := make(map[string]struct{})
	for _, v := range sl {
		myMap[v] = struct{}{}
	}
	fmt.Println(myMap)
}
