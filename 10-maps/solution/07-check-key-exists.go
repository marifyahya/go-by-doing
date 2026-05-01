package main

import "fmt"

func main() {
	m := map[string]int{"alice": 25}
	_, exists := m["alice"]
	fmt.Printf("alice exists: %t\n", exists)
	_, exists = m["charlie"]
	fmt.Printf("charlie exists: %t\n", exists)
}
