package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		<-ctx.Done()
		fmt.Println("Context done")
	}(ctx)
	cancel()
	time.Sleep(10 * time.Millisecond)
}
