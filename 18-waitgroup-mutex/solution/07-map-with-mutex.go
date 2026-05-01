package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	m := make(map[string]int)
	
	mu.Lock()
	m["key"] = 1
	mu.Unlock()
	
	fmt.Println("Safe map operations")
}
