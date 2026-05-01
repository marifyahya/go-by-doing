package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := &Person{Name: "Alice", Age: 25}
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Age: %d\n", p.Age)
}
