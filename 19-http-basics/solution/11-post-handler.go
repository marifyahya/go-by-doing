package main

import (
	"fmt"
	"net/http"
)

func main() {
	r, _ := http.NewRequest("POST", "/", nil)
	if r.Method == "POST" {
		fmt.Println("Data processed")
	}
}
