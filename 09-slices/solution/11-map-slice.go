package main

import "fmt"

func main() {
	original := []int{1, 2, 3}
	doubled := make([]int, len(original))
	for i, v := range original {
		doubled[i] = v * 2
	}
	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Doubled: %v\n", doubled)
}
