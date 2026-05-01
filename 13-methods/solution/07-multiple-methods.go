package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) GetName() string { return p.Name }
func (p Person) GetAge() int    { return p.Age }

func main() {
	p := Person{Name: "Alice", Age: 25}
	fmt.Println(p.GetName())
	fmt.Println(p.GetAge())
}
