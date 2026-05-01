package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

type Config struct {
	Message string
}

func configMiddleware(cfg Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(cfg.Message)
			next.ServeHTTP(w, r)
		})
	}
}

func main() {
	cfg := Config{Message: "Custom config used"}
	mw := configMiddleware(cfg)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	mw(handler).ServeHTTP(rr, req)
}
