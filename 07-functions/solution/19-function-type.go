package main

import "fmt"

type Operation func(int, int) int

func main() {
	var op Operation = func(a, b int) int {
		return a + b
	}
	fmt.Printf("Operation result: %d\n", op(5, 3))
}
