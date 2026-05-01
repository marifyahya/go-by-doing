package main

import "fmt"

func main() {
	dayNum := 3
	var dayName string

	switch dayNum {
	case 1:
		dayName = "Monday"
	case 2:
		dayName = "Tuesday"
	case 3:
		dayName = "Wednesday"
	case 4:
		dayName = "Thursday"
	case 5:
		dayName = "Friday"
	case 6:
		dayName = "Saturday"
	case 7:
		dayName = "Sunday"
	default:
		dayName = "Unknown"
	}

	fmt.Printf("%d: %s\n", dayNum, dayName)
}
