package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	mw := corsMiddleware(handler)
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	mw.ServeHTTP(rr, req)
	fmt.Printf("Access-Control-Allow-Origin: %s\n", rr.Header().Get("Access-Control-Allow-Origin"))
}
