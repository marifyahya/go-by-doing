package main

import (
	"fmt"
	"net/http"
)

func main() {
	r, _ := http.NewRequest("GET", "/?name=John&age=25", nil)
	query := r.URL.Query()
	fmt.Printf("Name: %s\n", query.Get("name"))
	fmt.Printf("Age: %s\n", query.Get("age"))
}
