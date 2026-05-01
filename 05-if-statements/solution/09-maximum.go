package main

import "fmt"

func main() {
	a, b := 10, 15

	// Find the maximum
	max := a
	if b > a {
		max = b
	}
	fmt.Printf("Max: %d\n", max)
}
