package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	data := []byte(`{"name":"Alice"}`)
	var u User
	json.Unmarshal(data, &u)
	fmt.Println(u.Name)
}
