package main

import "fmt"

func main() {
	a, b, c := 10, 5, 3

	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}

	fmt.Printf("Min: %d\n", min)
}
