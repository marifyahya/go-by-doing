package main

import (
	"fmt"
	"net/http"
)

func main() {
	r, _ := http.NewRequest("GET", "/", nil)
	fmt.Printf("Method: %s\n", r.Method)
}
