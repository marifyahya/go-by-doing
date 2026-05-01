package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[string]int)
	var mu sync.Mutex
	mu.Lock()
	m["key"] = 1
	mu.Unlock()
	fmt.Println("Safe map operations")
}
