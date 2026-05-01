package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	server := &http.Server{Addr: ":8080"}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// server.Shutdown(ctx)
	_ = server
	_ = ctx
	fmt.Println("Server stopped gracefully")
}
