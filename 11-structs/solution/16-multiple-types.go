package main

import "fmt"

type Person struct {
	Name string
}

type Address struct {
	City string
}

func main() {
	p := Person{Name: "Alice"}
	a := Address{City: "NYC"}
	fmt.Printf("Person: %s\n", p.Name)
	fmt.Printf("Address: %s\n", a.City)
}
