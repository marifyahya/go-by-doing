package main

import "fmt"

func main() {
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	total := 0
	for _, row := range matrix {
		for _, v := range row {
			total += v
		}
	}
	fmt.Printf("Total: %d\n", total)
}
