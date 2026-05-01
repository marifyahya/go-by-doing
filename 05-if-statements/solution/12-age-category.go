package main

import "fmt"

func main() {
	age := 25
	var category string

	if age < 13 {
		category = "Child"
	} else if age <= 19 {
		category = "Teen"
	} else if age <= 64 {
		category = "Adult"
	} else {
		category = "Senior"
	}

	fmt.Printf("Age: %d, Category: %s\n", age, category)
}
