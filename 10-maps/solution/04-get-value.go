package main

import "fmt"

func main() {
	m := map[string]int{"alice": 25, "bob": 30}
	
	val, ok := m["alice"]
	if ok {
		fmt.Printf("Alice: %d\n", val)
	} else {
		fmt.Println("Alice: (not found)")
	}

	val, ok = m["charlie"]
	if ok {
		fmt.Printf("Charlie: %d\n", val)
	} else {
		fmt.Println("Charlie: (not found)")
	}
}
