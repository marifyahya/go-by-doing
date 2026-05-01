package main

import (
	"encoding/json"
	"fmt"
)

type Profile struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type User struct {
	Profile Profile `json:"user"`
}

func main() {
	u := User{Profile: Profile{Name: "Alice", Email: "alice@test.com"}}
	data, _ := json.Marshal(u)
	fmt.Println(string(data))
}
