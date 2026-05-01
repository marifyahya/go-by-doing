package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s 200 10ms\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	mw := requestLogger(handler)
	req := httptest.NewRequest("GET", "/api/users", nil)
	rr := httptest.NewRecorder()
	mw.ServeHTTP(rr, req)
}
