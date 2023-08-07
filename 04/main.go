package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(wg *sync.WaitGroup, ch <-chan string, id int) {
	defer wg.Done()
	for v := range ch {
		time.Sleep(time.Millisecond)
		fmt.Printf("[worker %4d]: %s\n", id, v)
	}
	fmt.Printf("[Worker %4d]: Done\n", id)
}

func main() {

	var nums int
	ch := make(chan string)
	fmt.Print("Enter number of workers: ")
	fmt.Scan(&nums)
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	//graceful shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig
		fmt.Println("Gracefully shutdown")
		cancel()
		time.Sleep(time.Second)
		close(ch)
		wg.Wait()
		os.Exit(0)
	}()
	for i := 0; i < nums; i++ {
		wg.Add(1)
		go worker(wg, ch, i)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		var value string
		for {
			fmt.Println("Enter anything to chan")
			fmt.Scan(&value)
			select {
			case <-ctx.Done():
				return
			default:
				ch <- value
			}

		}
	}()
	wg.Wait()
	fmt.Println("END")
}
