package main

import "fmt"

func transform(nums []int, f func(int) int) []int {
	res := make([]int, len(nums))
	for i, v := range nums {
		res[i] = f(v)
	}
	return res
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	res := transform(nums, func(x int) int {
		return x * 2
	})
	fmt.Printf("Transformed: %v\n", res)
}
