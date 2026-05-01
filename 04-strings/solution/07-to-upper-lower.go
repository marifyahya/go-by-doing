package main

import "fmt"
import "strings"

func main() {
	s := "hello world"
	fmt.Printf("Upper: %s\n", strings.ToUpper(s))
	fmt.Printf("Lower: %s\n", strings.ToLower(s))
	fmt.Printf("Title: %s\n", strings.ToTitle(s))
}