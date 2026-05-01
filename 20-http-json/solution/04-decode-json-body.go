package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	body := strings.NewReader(`{"name":"Alice"}`)
	var u User
	json.NewDecoder(body).Decode(&u)
	fmt.Printf("Decoded: %s\n", u.Name)
}
