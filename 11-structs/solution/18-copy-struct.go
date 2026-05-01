package main

import "fmt"

type Person struct {
	Name string
}

func main() {
	p1 := Person{Name: "Alice"}
	p2 := p1 // structs are values, this copies
	fmt.Printf("Original: %s\n", p1.Name)
	fmt.Printf("Copy: %s\n", p2.Name)
}
