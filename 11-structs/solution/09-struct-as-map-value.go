package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {
	m := map[string]Person{
		"Alice": {Name: "Alice", Age: 25, Email: "alice@test.com"},
	}
	fmt.Printf("Alice: %+v\n", m["Alice"])
}
