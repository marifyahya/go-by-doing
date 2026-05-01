package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() { ch1 <- 1 }()

	select {
	case <-ch1:
		fmt.Println("Received from one of the channels")
	case <-ch2:
		fmt.Println("Received from one of the channels")
	}
}
