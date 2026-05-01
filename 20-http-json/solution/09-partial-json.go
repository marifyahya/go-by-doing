package main

import (
	"encoding/json"
	"fmt"
)

type PartialUser struct {
	Name string `json:"name"`
}

func main() {
	data := `{"name":"Alice","age":25,"email":"alice@test.com"}`
	var u PartialUser
	json.Unmarshal([]byte(data), &u)
	fmt.Println("Only name decoded")
}
