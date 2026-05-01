package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	select {
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Done")
	case <-ctx.Done():
		fmt.Println("Timed out")
	}
}
