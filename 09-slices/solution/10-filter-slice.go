package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	var evens []int
	for _, v := range slice {
		if v%2 == 0 {
			evens = append(evens, v)
		}
	}
	fmt.Printf("Original: %v\n", slice)
	fmt.Printf("Evens: %v\n", evens)
}
