package main

import (
	"fmt"
	"time"
)

func main() {
	go func() { fmt.Println("Task 1 done") }()
	go func() { fmt.Println("Task 2 done") }()
	time.Sleep(10 * time.Millisecond)
}
