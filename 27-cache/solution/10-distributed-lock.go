package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func main() {
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	_ = rdb.SetNX(context.Background(), "mylock", "locked", 10*time.Second)
	fmt.Println("Lock acquired")
}
