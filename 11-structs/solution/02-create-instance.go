package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Alice", Age: 25}
	fmt.Printf("%s, %d\n", p.Name, p.Age)
}
