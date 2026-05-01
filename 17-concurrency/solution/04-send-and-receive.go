package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		val := 10
		fmt.Printf("Sent: %d\n", val)
		ch <- val
	}()
	fmt.Printf("Received: %d\n", <-ch)
}
