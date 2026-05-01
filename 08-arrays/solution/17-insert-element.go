package main

import "fmt"

func main() {
	// Arrays in Go are fixed size. To "insert", we need a larger array.
	original := [5]int{1, 2, 3, 4, 5}
	var inserted [6]int
	pos := 2
	val := 99

	for i := 0; i < pos; i++ {
		inserted[i] = original[i]
	}
	inserted[pos] = val
	for i := pos; i < len(original); i++ {
		inserted[i+1] = original[i]
	}
	fmt.Println(inserted)
}
