package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	u := User{Name: "Alice", Age: 25}
	data, _ := json.Marshal(u)
	fmt.Println(string(data))
}
