package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func basicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("Middleware -> ")
		next.ServeHTTP(w, r)
	})
}

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler")
	})

	mw := basicMiddleware(handler)
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	mw.ServeHTTP(rr, req)
}
