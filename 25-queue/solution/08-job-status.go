package main

import (
	"fmt"
)

func main() {
	// status, _ := rdb.Get(ctx, "job:1:status").Result()
	status := "completed"
	fmt.Printf("Status: %s\n", status)
}
