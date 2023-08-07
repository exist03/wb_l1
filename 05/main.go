package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func f(t *time.Time) int {
	return int(time.Since(*t).Seconds())
}

func main() {
	var input string
	fmt.Print("Enter how many program should work: ")
	fmt.Scan(&input)
	ch := make(chan int)
	duration, _ := time.ParseDuration(input)
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n := rand.Int()
				ch <- n
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case v := <-ch:
				fmt.Println(v)
			case <-ctx.Done():
				return
			}
		}
	}()

	wg.Wait()
	close(ch)
}
