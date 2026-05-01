package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()

	fmt.Print("Values:")
	for v := range ch {
		fmt.Printf(" %d", v)
	}
	fmt.Println()
	fmt.Println("Closed")
}
