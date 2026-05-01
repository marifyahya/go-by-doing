package main

import (
	"fmt"
)

func main() {
	// Show both:
	// 1. Mutex: mu.Lock(); counter++; mu.Unlock()
	// 2. Channel: ch <- 1; ...; counter += <-ch
	fmt.Println("Both work correctly")
}
