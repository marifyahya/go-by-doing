package main

import "fmt"

func main() {
	m := map[string]map[string]string{
		"Alice": {"city": "New York"},
	}
	fmt.Printf("Alice's city: %s\n", m["Alice"]["city"])
}
