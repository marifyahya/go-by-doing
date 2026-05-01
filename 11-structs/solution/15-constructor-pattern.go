package main

import "fmt"

type Person struct {
	Name string
}

func NewPerson(name string) *Person {
	return &Person{Name: name}
}

func main() {
	p := NewPerson("Alice")
	fmt.Printf("Created: %s\n", p.Name)
}
