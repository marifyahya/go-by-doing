package main

import (
	"fmt"
	"net/http/httptest"
)

func main() {
	req := httptest.NewRequest("GET", "/users?name=john", nil)
	fmt.Printf("Path: %s\n", req.URL.Path)
	fmt.Printf("Query: %s\n", req.URL.RawQuery)
}
