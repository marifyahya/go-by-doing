package main

import "fmt"

func main() {
	original := [5]int{1, 2, 3, 4, 5}
	var deleted [4]int
	pos := 2

	for i := 0; i < pos; i++ {
		deleted[i] = original[i]
	}
	for i := pos + 1; i < len(original); i++ {
		deleted[i-1] = original[i]
	}
	fmt.Println(deleted)
}
