package main

import "fmt"

func main() {
	score := 95
	attendance := 90

	// Nested If example
	if score >= 75 {
		if attendance >= 80 {
			fmt.Println("Pass with distinction")
		} else {
			fmt.Println("Pass")
		}
	} else {
		fmt.Println("Fail")
	}
}
