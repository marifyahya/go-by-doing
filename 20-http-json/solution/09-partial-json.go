package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	data := []byte(`{"name":"Alice","age":25}`)
	var u User
	json.Unmarshal(data, &u)
	fmt.Println("Only name decoded")
}
