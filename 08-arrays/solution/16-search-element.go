package main

import "fmt"

func main() {
	arr := [5]int{10, 20, 30, 40, 50}
	target := 30
	foundIndex := -1
	for i, v := range arr {
		if v == target {
			foundIndex = i
			break
		}
	}
	fmt.Printf("Found at index: %d\n", foundIndex)

	target = 99
	foundIndex = -1
	for i, v := range arr {
		if v == target {
			foundIndex = i
			break
		}
	}
	fmt.Printf("Not found: %d\n", foundIndex)
}
