package main

import "fmt"

func main() {
	score := 75
	if score >= 90 {
		fmt.Println("Excellent")
	} else if score >= 70 {
		fmt.Println("Good")
	} else {
		fmt.Println("Needs improvement")
	}
}