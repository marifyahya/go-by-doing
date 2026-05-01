package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("(background task runs)")
	}()
	fmt.Println("Main done")
	time.Sleep(10 * time.Millisecond)
}
