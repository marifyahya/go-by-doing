package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Printf("Sum: %d\n", add(7, 8))
}
