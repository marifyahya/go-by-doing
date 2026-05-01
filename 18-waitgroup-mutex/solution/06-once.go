package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Initialized once")
	}
	for i := 0; i < 10; i++ {
		once.Do(onceBody)
	}
}
