package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	data := `{"name":"Alice"}`
	var u User
	json.Unmarshal([]byte(data), &u)
	fmt.Println(u.Name)
}
