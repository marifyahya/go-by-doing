package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	// val, _ := rdb.Get(context.Background(), "user").Result()
	val := "user123"
	fmt.Printf("Value: %s\n", val)
}
