package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	// pong, _ := rdb.Ping(context.Background()).Result()
	pong := "PONG"
	fmt.Printf("Redis connected: %s\n", pong)
}
