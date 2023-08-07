package main

import "fmt"

func main() {
	arr := []int{1, 5, 1, 36, 8, 9, 212, -12}
	qsort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
func qsort(arr []int, start, end int) {
	if start >= end {
		return
	}
	pivot := partition(arr, start, end)
	qsort(arr, start, pivot-1)
	qsort(arr, pivot+1, end)

}
func partition(arr []int, start, end int) int {
	i := start - 1
	for j := start; j < end; j++ {
		if arr[j] < arr[end] {
			i++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	i++
	arr[i], arr[end] = arr[end], arr[i]
	return i
}
