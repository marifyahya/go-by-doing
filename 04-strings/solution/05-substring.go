package main

import "fmt"

func main() {
	s := "Hello World"
	fmt.Printf("First 5: %s\n", s[:5])
	fmt.Printf("From 6: %s\n", s[6:])
	fmt.Printf("Middle: %s\n", s[3:8])
}