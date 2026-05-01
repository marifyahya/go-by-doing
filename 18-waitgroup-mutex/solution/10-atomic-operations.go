package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var count int64
	for i := 0; i < 1000; i++ {
		atomic.AddInt64(&count, 1)
	}
	fmt.Printf("Value: %d\n", atomic.LoadInt64(&count))
}
