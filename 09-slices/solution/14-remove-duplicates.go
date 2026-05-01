package main

import "fmt"

func main() {
	slice := []int{1, 2, 2, 3, 4, 4, 5}
	uniqueMap := make(map[int]bool)
	var uniqueSlice []int
	for _, v := range slice {
		if !uniqueMap[v] {
			uniqueMap[v] = true
			uniqueSlice = append(uniqueSlice, v)
		}
	}
	fmt.Println(uniqueSlice)
}
