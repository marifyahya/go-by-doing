package main

import "fmt"
import "strings"

func main() {
	s := "one two three"
	parts := strings.Split(s, " ")
	fmt.Printf("Parts: %v\n", parts)
	fmt.Printf("Count: %d\n", len(parts))
}