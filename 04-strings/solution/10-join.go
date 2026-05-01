package main

import "fmt"
import "strings"

func main() {
	parts := []string{"one", "two", "three"}
	joined := strings.Join(parts, "-")
	fmt.Printf("Joined: %s\n", joined)
}