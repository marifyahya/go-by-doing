package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "userID", 123)
	fmt.Printf("UserID: %v\n", ctx.Value("userID"))
}
