package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	wg := &sync.WaitGroup{}
	//closed chanel
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case _, ok := <-ch:
				if !ok {
					fmt.Println("chanel is closed")
					return
				}
			}
		}
	}()
	ch <- 1
	ch <- 2
	close(ch)
	wg.Wait()

	//special chanel(it can be chan bool or ctx.Done())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-sig
		fmt.Println("another method")
		return
	}()
	wg.Wait()
}
