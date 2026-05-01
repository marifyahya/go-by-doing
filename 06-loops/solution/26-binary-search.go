package main

import "fmt"

func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 5, 8, 13, 21}
	target := 8
	index := binarySearch(arr, target)
	fmt.Printf("Found at index: %d\n", index)
}
