package main

import "fmt"

func main() {
	nums := []int{3, 7, 2, 9, 1, 5}
	min := nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	fmt.Printf("Min: %d\n", min)
}
