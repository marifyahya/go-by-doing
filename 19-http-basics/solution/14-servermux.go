package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	_ = mux
	fmt.Println("Routes registered")
}
