package main

import (
	"fmt"
	"net/http"
)

func main() {
	r, _ := http.NewRequest("GET", "/", nil)
	r.Header.Set("Content-Type", "application/json")
	fmt.Printf("Content-Type: %s\n", r.Header.Get("Content-Type"))
}
