package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) Info() string {
	return "Person: " + p.Name
}

type Employee struct {
	Person
	Company string
}

func (e Employee) Info() string {
	return fmt.Sprintf("Employee: %s, %s", e.Name, e.Company)
}

func main() {
	e := Employee{Person{Name: "Alice"}, "TechCorp"}
	fmt.Println(e.Info())
}
