package main

import "fmt"

func main() {
	var (
		name string = "Alice"
		age  int    = 25
	)
	fmt.Println(name, age)

	var (
		name2 string = "Bob"
		age2  int    = 30
	)
	fmt.Println(name2, age2)
}