package main

import "fmt"

func main() {
	nums := []int{15, 10}

	for _, n := range nums {
		if n%3 == 0 && n%5 == 0 {
			fmt.Printf("%d: Divisible by both\n", n)
		} else if n%3 == 0 {
			fmt.Printf("%d: Divisible by 3 only\n", n)
		} else if n%5 == 0 {
			fmt.Printf("%d: Divisible by 5 only\n", n)
		}
	}
}
