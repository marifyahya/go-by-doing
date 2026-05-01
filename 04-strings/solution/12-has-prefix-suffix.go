package main

import "fmt"
import "strings"

func main() {
	s := "Hello World"
	fmt.Printf("HasPrefix: %t\n", strings.HasPrefix(s, "Hello"))
	fmt.Printf("HasSuffix: %t\n", strings.HasSuffix(s, "World"))
}