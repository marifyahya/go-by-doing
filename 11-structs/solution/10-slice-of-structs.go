package main

import "fmt"

type Person struct {
	Name string
}

func main() {
	people := []Person{
		{Name: "Alice"},
		{Name: "Bob"},
		{Name: "Charlie"},
	}
	var names []string
	for _, p := range people {
		names = append(names, p.Name)
	}
	fmt.Println(names)
}
