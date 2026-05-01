package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for range jobs {
		fmt.Printf("Worker %d processed job\n", id)
	}
}

func main() {
	jobs := make(chan int, 2)
	var wg sync.WaitGroup

	for w := 1; w <= 2; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	jobs <- 1
	jobs <- 2
	close(jobs)
	wg.Wait()
}
