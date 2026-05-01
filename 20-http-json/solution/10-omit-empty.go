package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
}

func main() {
	u := User{Name: "Alice"}
	data, _ := json.Marshal(u)
	fmt.Println(string(data))
}
