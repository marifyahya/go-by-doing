package main

import "fmt"

func main() {
	name := "Alice"
	age := 25
	s := fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Println(s)
}