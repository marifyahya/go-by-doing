package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	r, _ := http.NewRequest("GET", "/users/123", nil)
	parts := strings.Split(r.URL.Path, "/")
	fmt.Printf("User ID: %s\n", parts[2])
}
