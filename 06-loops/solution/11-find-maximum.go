package main

import "fmt"

func main() {
	nums := []int{3, 7, 2, 9, 5}
	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	fmt.Printf("Max: %d\n", max)
}
