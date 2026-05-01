package main

import "fmt"

func main() {
	source := []int{1, 2, 3}
	copySlice := make([]int, len(source))
	copy(copySlice, source)
	fmt.Printf("Source: %v\n", source)
	fmt.Printf("Copy: %v\n", copySlice)
}
