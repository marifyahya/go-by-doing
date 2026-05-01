package main

import "fmt"

type Person struct {
	Name string
}

type Employee struct {
	Person
	Company string
}

func main() {
	e := Employee{
		Person:  Person{Name: "Alice"},
		Company: "TechCorp",
	}
	fmt.Printf("Employee: %s, Company: %s\n", e.Name, e.Company)
}
