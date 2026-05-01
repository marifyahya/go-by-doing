package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
)

func contextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "userID", "123")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("userID")
		fmt.Printf("User ID: %v\n", userID)
	})
	mw := contextMiddleware(handler)
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	mw.ServeHTTP(rr, req)
}
