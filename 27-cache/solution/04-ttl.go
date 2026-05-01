package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func main() {
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	_ = rdb.Set(context.Background(), "key", "value", 1*time.Hour)
	fmt.Println("Expires in 1 hour")
}
