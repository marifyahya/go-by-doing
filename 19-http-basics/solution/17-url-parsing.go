package main

import (
	"fmt"
	"net/http"
)

func main() {
	r, _ := http.NewRequest("GET", "/users?name=john", nil)
	fmt.Printf("Path: %s\n", r.URL.Path)
	fmt.Printf("Query: %s\n", r.URL.RawQuery)
}
