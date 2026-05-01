package main

import "fmt"

func main() {
	arr := [5]int{3, 1, 4, 1, 5}
	fmt.Printf("Unsorted: %v\n", arr)

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Printf("Sorted: %v\n", arr)
}
