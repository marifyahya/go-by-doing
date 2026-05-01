package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

func main() {
	req := httptest.NewRequest("GET", "/users/123", nil)
	parts := strings.Split(req.URL.Path, "/")
	if len(parts) > 2 {
		fmt.Printf("User ID: %s\n", parts[2])
	}
}
