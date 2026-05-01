package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func main() {
	w := httptest.NewRecorder()
	http.NotFound(w, nil)
	fmt.Println("404 Not Found")
}
