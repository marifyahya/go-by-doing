package main

import "fmt"

type Person struct {
	Name string
	ID   int
}

func (p Person) Info() string {
	return fmt.Sprintf("%s (ID: %d)", p.Name, p.ID)
}

func main() {
	p := Person{Name: "Alice", ID: 1}
	fmt.Println(p.Info())
}
