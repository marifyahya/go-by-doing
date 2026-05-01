package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	go func() {
		time.Sleep(10 * time.Millisecond)
		c.Signal()
	}()
	c.L.Lock()
	c.Wait()
	c.L.Unlock()
	fmt.Println("Signal received")
}
