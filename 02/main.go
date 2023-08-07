package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 8, 10}
	wg := &sync.WaitGroup{}
	for i := range arr {
		wg.Add(1)
		go func(value int) {
			defer wg.Done()
			fmt.Println(value)
		}(arr[i])
	}
	wg.Wait()
}
