package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	_ = rdb.LPush(context.Background(), "jobs", `{"type": "email"}`)
	fmt.Println("Job queued")
}
