package main

import "fmt"

func main() {
	num := 10

	// Check if positive, negative, or zero
	if num > 0 {
		fmt.Println("Positive")
	} else if num < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Zero")
	}
}
