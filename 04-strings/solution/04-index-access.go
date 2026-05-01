package main

import "fmt"

func main() {
	s := "Hello"
	fmt.Printf("First char: %c\n", s[0])
	fmt.Printf("Last char: %c\n", s[len(s)-1])
}