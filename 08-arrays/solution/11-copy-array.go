package main

import "fmt"

func main() {
	original := [3]int{1, 2, 3}
	var copyArr [3]int
	for i, v := range original {
		copyArr[i] = v
	}
	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Copy: %v\n", copyArr)
}
