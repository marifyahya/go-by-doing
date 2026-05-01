package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func main() {
	req := httptest.NewRequest("POST", "/", nil)
	if req.Method == "POST" {
		fmt.Println("Data processed")
	}
}
