package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("Cancelled")
		}
	}()
	cancel()
	time.Sleep(10 * time.Millisecond)
}
