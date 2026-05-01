package main

import "fmt"

func main() {
	arr := [5]int{3, 7, 2, 9, 5}
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	fmt.Printf("Max: %d\n", max)
}
