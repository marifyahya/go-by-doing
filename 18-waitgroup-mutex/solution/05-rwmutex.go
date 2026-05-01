package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.RWMutex
	// Simulate many reads
	mu.RLock()
	fmt.Println("Read operations many")
	mu.RUnlock()

	// Simulate few writes
	mu.Lock()
	fmt.Println("Write operations few")
	mu.Unlock()
}
