package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func main() {
	rr := httptest.NewRecorder()
	rr.WriteHeader(http.StatusOK)
	fmt.Println("Status: 200 OK")
}
