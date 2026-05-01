package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func rateLimitMiddleware(next http.Handler) http.Handler {
	count := 0
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		count++
		if count > 0 { // Simulating immediate limit for demo
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			fmt.Println("429 Too Many Requests")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	mw := rateLimitMiddleware(handler)
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	mw.ServeHTTP(rr, req)
}
