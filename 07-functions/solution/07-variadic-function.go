package main

import "fmt"

func sumAll(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	fmt.Printf("Sum: %d\n", sumAll(1, 2, 3, 4, 5))
	fmt.Printf("Sum: %d\n", sumAll(10, 10))
}
