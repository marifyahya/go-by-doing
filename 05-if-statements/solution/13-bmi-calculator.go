package main

import "fmt"

func main() {
	weight := 70.0
	height := 1.75
	bmi := weight / (height * height)
	var category string

	if bmi < 18.5 {
		category = "Underweight"
	} else if bmi <= 25 {
		category = "Normal"
	} else if bmi <= 30 {
		category = "Overweight"
	} else {
		category = "Obese"
	}

	fmt.Printf("Weight: %.0f, Height: %.2f, BMI: %.2f, Category: %s\n", weight, height, bmi, category)
}
