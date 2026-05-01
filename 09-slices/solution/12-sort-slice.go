package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{3, 1, 4, 1, 5}
	fmt.Printf("Unsorted: %v\n", slice)
	sort.Ints(slice)
	fmt.Printf("Sorted: %v\n", slice)
}
