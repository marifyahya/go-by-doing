package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	r, _ := http.NewRequest("POST", "/", nil)
	r.Form = url.Values{"username": {"john"}}
	fmt.Printf("Username: %s\n", r.Form.Get("username"))
}
