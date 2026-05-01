package main

import "fmt"

func main() {
	x := 5
	fmt.Printf("Initial: %d\n", x)

	x++
	fmt.Printf("After ++: %d\n", x)

	x--
	fmt.Printf("After --: %d\n", x)
}