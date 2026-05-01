package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func main() {
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Content-Type", "application/json")
	fmt.Printf("Content-Type: %s\n", req.Header.Get("Content-Type"))
}
