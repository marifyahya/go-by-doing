package main

import "fmt"

func main() {
	person := struct {
		Name string
		Age  int
	}{
		Name: "John",
		Age:  30,
	}
	fmt.Printf("%s, %d\n", person.Name, person.Age)
}
