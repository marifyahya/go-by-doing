package main

import "fmt"

type Address struct {
	City    string
	Country string
}

type Person struct {
	Name    string
	Address Address
}

func main() {
	p := Person{
		Name: "Alice",
		Address: Address{
			City:    "New York",
			Country: "USA",
		},
	}
	fmt.Printf("%s lives in %s, %s\n", p.Name, p.Address.City, p.Address.Country)
}
