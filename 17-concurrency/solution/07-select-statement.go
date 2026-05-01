package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go func() {
		time.Sleep(10 * time.Millisecond)
		ch1 <- 1
	}()

	select {
	case v := <-ch1:
		fmt.Printf("Received from ch1: %d\n", v)
	}
}
