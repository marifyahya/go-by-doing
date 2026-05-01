package main

import "fmt"

func main() {
	num := -5

	if num > 0 {
		fmt.Printf("%d is positive\n", num)
	} else if num < 0 {
		fmt.Printf("%d is negative\n", num)
	} else {
		fmt.Printf("%d is zero\n", num)
	}
}
