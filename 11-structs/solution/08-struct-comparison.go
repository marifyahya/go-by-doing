package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p1 := Person{Name: "Alice", Age: 25}
	p2 := Person{Name: "Alice", Age: 25}
	fmt.Printf("Equal: %t\n", p1 == p2)
}
