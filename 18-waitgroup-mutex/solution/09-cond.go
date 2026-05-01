package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mu := &sync.Mutex{}
	cond := sync.NewCond(mu)
	
	go func() {
		time.Sleep(10 * time.Millisecond)
		cond.Signal()
	}()
	
	mu.Lock()
	cond.Wait()
	mu.Unlock()
	fmt.Println("Signal received")
}
