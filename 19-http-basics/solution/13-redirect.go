package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func main() {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/old", nil)
	http.Redirect(w, req, "/new", http.StatusMovedPermanently)
	fmt.Println("Redirected to /new")
}
