package main

import "fmt"

type User struct {
	Name  string
	Email string
}

func main() {
	m := map[string]User{
		"alice": {Name: "Alice", Email: "alice@test.com"},
	}
	u := m["alice"]
	fmt.Printf("User: %s, Email: %s\n", u.Name, u.Email)
}
