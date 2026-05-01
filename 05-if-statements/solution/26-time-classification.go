package main

import "fmt"

func main() {
	hour := 14

	if hour >= 5 && hour <= 11 {
		fmt.Println("Morning")
	} else if hour >= 12 && hour <= 17 {
		fmt.Println("Afternoon")
	} else if hour >= 18 && hour <= 21 {
		fmt.Println("Evening")
	} else {
		fmt.Println("Night")
	}
}
