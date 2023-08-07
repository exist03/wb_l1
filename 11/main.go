package main

import "fmt"

func main() {
	arr1 := []int{1, 6, 4, 8, 5, 2}
	arr2 := []int{0, -5, 3, 5, 6, 2}
	fmt.Println(f(arr1, arr2))
}

func f(arr1, arr2 []int) []int {
	myMap := make(map[int]int)
	result := make([]int, 0)
	if len(arr1) > len(arr2) {
		arr1, arr2 = arr2, arr1
	}
	for _, v := range arr1 {
		myMap[v] = 0
	}
	for _, v := range arr2 {
		if _, ok := myMap[v]; ok {
			result = append(result, v)
		}
	}
	return result
}
