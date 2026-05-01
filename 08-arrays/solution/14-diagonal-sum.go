package main

import "fmt"

func main() {
	matrix := [3][3]int{
		{5, 0, 0},
		{0, 5, 0},
		{0, 0, 5},
	}
	sum := 0
	for i := 0; i < 3; i++ {
		sum += matrix[i][i]
	}
	fmt.Printf("Diagonal sum: %d\n", sum)
}
