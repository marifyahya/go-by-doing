package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

func main() {
	form := "username=john"
	req := httptest.NewRequest("POST", "/", strings.NewReader(form))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.ParseForm()
	fmt.Printf("Username: %s\n", req.Form.Get("username"))
}
