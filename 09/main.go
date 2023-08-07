package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{0, 1, 2, 3, 4, 5}
	firstCh := make(chan int, len(arr))
	secondCh := make(chan int, len(arr))
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(firstCh)
		for _, v := range arr {
			firstCh <- v
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(secondCh)
		for _, v := range arr {
			secondCh <- v * 2
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range secondCh {
			fmt.Println(v)
		}
	}()
	wg.Wait()
}
