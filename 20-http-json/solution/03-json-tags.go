package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"user_name"`
	Age  int    `json:"user_age"`
}

func main() {
	u := User{Name: "alice", Age: 25}
	data, _ := json.Marshal(u)
	fmt.Println(string(data))
}
