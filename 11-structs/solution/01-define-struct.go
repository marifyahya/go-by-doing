package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var p Person
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}
