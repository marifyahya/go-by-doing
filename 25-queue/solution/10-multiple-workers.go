package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d processing\n", id)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(1, &wg)
	go worker(2, &wg)
	wg.Wait()
}
