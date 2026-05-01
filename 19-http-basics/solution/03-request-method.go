package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func main() {
	req := httptest.NewRequest("GET", "/", nil)
	fmt.Printf("Method: %s\n", req.Method)
}
