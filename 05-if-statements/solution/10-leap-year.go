package main

import "fmt"

func main() {
	year1 := 2024
	year2 := 2023

	isLeap := func(y int) bool {
		return y%4 == 0 && (y%100 != 0 || y%400 == 0)
	}

	fmt.Printf("%d is a leap year: %t\n", year1, isLeap(year1))
	fmt.Printf("%d is a leap year: %t\n", year2, isLeap(year2))
}
