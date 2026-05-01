package main

import "fmt"

type Person struct {
	Name string
}

func getPerson(name string) *Person {
	return &Person{Name: name}
}

func main() {
	p := getPerson("Alice")
	fmt.Println(p.Name)
}
