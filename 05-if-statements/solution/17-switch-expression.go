package main

import "fmt"

func main() {
	day := "Saturday"

	switch {
	case day == "Saturday" || day == "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}
}
