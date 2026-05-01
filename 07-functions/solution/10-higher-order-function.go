package main

import "fmt"

func apply(nums []int, f func(int) int) int {
	total := 0
	for _, n := range nums {
		total += f(n)
	}
	return total
}

func main() {
	nums := []int{1, 2, 3, 4}
	res := apply(nums, func(x int) int {
		return x * x
	})
	fmt.Printf("Result: %d\n", res)
}
