package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func main() {
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	// rdb.ZAdd(ctx, "delayed_jobs", redis.Z{Score: float64(time.Now().Add(1*time.Hour).Unix()), Member: "job1"})
	fmt.Println("Job scheduled for later")
}
