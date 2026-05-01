package main

import "fmt"

type Person struct {
	FirstName, LastName string
}

func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

func main() {
	p := Person{FirstName: "Alice", LastName: "Smith"}
	fmt.Printf("Full Name: %s\n", p.FullName())
}
