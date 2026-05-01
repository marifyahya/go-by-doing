package main

import "fmt"

func main() {
	age1 := 17
	age2 := 18

	check := func(age int) {
		if age >= 18 {
			fmt.Printf("Age %d: Eligible\n", age)
		} else {
			fmt.Printf("Age %d: Not eligible\n", age)
		}
	}

	check(age1)
	check(age2)
}
