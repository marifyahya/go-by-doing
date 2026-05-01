package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) Person {
	return Person{Name: name, Age: age}
}

func main() {
	p := NewPerson("Alice", 25)
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}
