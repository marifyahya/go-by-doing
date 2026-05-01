package main

import "fmt"

func main() {
	input := []string{"a", "b", "a", "c", "b"}
	fmt.Printf("Input: %v\n", input)
	
	uniqueMap := make(map[string]bool)
	var unique []string
	for _, s := range input {
		if !uniqueMap[s] {
			uniqueMap[s] = true
			unique = append(unique, s)
		}
	}
	fmt.Printf("Unique: %v\n", unique)
}
