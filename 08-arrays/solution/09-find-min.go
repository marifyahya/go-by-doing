package main

import "fmt"

func main() {
	arr := [5]int{3, 7, 2, 9, 1}
	min := arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	fmt.Printf("Min: %d\n", min)
}
