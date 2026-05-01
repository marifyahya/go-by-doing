package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	var wg sync.WaitGroup

	wg.Add(2)
	// To avoid deadlock, always acquire resources in the same order
	go func() {
		defer wg.Done()
		ch1 <- 1
		ch2 <- 2
	}()

	go func() {
		defer wg.Done()
		<-ch1
		<-ch2
	}()

	wg.Wait()
	fmt.Println("No deadlock")
}
