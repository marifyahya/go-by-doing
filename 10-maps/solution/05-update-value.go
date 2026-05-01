package main

import "fmt"

func main() {
	m := map[string]int{"alice": 25}
	fmt.Printf("Before: alice: %d\n", m["alice"])
	m["alice"] = 26
	fmt.Printf("After: alice: %d\n", m["alice"])
}
