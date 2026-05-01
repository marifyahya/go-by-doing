package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /todos", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `[]`)
	})
	fmt.Println("REST API Server running on :8080")
	// http.ListenAndServe(":8080", mux)
}
