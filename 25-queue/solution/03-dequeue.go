package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	// job, _ := rdb.BRPop(context.Background(), 0, "jobs").Result()
	fmt.Println("Job: {type: \"email\"}")
}
