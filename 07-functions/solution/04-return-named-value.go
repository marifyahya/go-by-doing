package main

import "fmt"

func addNamed(a, b int) (result int) {
	result = a + b
	return
}

func main() {
	fmt.Printf("Result: %d\n", addNamed(7, 8))
}
