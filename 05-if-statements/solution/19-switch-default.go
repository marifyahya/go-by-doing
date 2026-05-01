package main

import "fmt"

func main() {
	day := "Octday"

	switch day {
	case "Monday":
		fmt.Println("Monday")
	case "Tuesday":
		fmt.Println("Tuesday")
	default:
		fmt.Println("Invalid day")
	}
}
