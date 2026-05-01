package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.RWMutex
	_ = mu
	fmt.Println("Read operations many")
	fmt.Println("Write operations few")
}
