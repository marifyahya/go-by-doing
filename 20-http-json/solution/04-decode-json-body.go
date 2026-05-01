package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	body := `{"name":"Alice"}`
	req := httptest.NewRequest("POST", "/", strings.NewReader(body))
	var u User
	json.NewDecoder(req.Body).Decode(&u)
	fmt.Printf("Decoded: %s\n", u.Name)
}
