package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	pubsub := rdb.Subscribe(context.Background(), "mychannel")
	defer pubsub.Close()
	// msg, _ := pubsub.ReceiveMessage(ctx)
	fmt.Println("Message received")
}
