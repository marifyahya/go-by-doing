package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func main() {
	req := httptest.NewRequest("GET", "/?name=John&age=25", nil)
	query := req.URL.Query()
	fmt.Printf("Name: %s\n", query.Get("name"))
	fmt.Printf("Age: %s\n", query.Get("age"))
}
