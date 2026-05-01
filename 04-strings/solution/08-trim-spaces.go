package main

import "fmt"
import "strings"

func main() {
	s := "   Hello World   "
	fmt.Printf("Trimmed: %s\n", strings.TrimSpace(s))
	fmt.Printf("TrimLeft: %s\n", strings.TrimLeft(s, " "))
	fmt.Printf("TrimRight: %s\n", strings.TrimRight(s, " "))
}