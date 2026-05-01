package main

import "fmt"
import "strings"

func main() {
	s := "Hello World"
	fmt.Printf("Original: %s\n", s)
	replaced := strings.Replace(s, "World", "Go", 1)
	fmt.Printf("Replaced: %s\n", replaced)
}