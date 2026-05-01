package main

import "fmt"

func main() {
	nested := [][]int{{1, 2}, {3, 4}, {5, 6}}
	var flattened []int
	for _, sub := range nested {
		flattened = append(flattened, sub...)
	}
	fmt.Println(flattened)
}
