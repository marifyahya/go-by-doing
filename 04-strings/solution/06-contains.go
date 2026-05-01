package main

import "fmt"
import "strings"

func main() {
	s := "Hello World"
	fmt.Printf("Contains \"ello\": %t\n", strings.Contains(s, "ello"))
	fmt.Printf("Contains \"xyz\": %t\n", strings.Contains(s, "xyz"))
}