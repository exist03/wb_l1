package main

import "fmt"

func main() {
	sl := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	a := binarySearch(sl, 4)
	fmt.Println(a)
}
func binarySearch(sl []int, target int) bool {
	l, r := 0, len(sl)-1
	for l <= r {
		if target > sl[(l+r)/2] {
			l = (l+r)/2 + 1
		} else if target < sl[(l+r)/2] {
			r = (l+r)/2 - 1
		} else {
			return true
		}
	}
	return false
}
