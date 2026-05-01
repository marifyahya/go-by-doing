package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	select {
	case <-ch:
		fmt.Println("Received")
	case <-time.After(10 * time.Millisecond):
		fmt.Println("Timeout!")
	}
}
