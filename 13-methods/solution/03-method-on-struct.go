package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Info() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

func main() {
	p := Person{Name: "Alice", Age: 25}
	fmt.Println(p.Info())
}
