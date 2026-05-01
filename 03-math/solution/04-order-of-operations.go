package main

import "fmt"

func main() {
	result1 := 2 + 3*4
	result2 := (2 + 3) * 4

	fmt.Printf("2 + 3 * 4 = %d\n", result1)
	fmt.Printf("(2 + 3) * 4 = %d\n", result2)
}