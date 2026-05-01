package main

import "fmt"

type Person struct {
	Age int
}

func main() {
	var p Person
	fmt.Printf("{ %d }\n", p.Age)
}
