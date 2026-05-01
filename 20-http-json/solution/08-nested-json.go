package main

import (
	"encoding/json"
	"fmt"
)

type Details struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type User struct {
	User Details `json:"user"`
}

func main() {
	u := User{User: Details{Name: "Alice", Email: "alice@test.com"}}
	data, _ := json.Marshal(u)
	fmt.Println(string(data))
}
