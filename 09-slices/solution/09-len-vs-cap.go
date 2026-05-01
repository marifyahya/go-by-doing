package main

import "fmt"

func main() {
	slice := make([]int, 3, 5)
	fmt.Printf("Length: %d, Capacity: %d\n", len(slice), cap(slice))
	slice = append(slice, 4, 5, 6)
	fmt.Printf("Length: %d, Capacity: %d\n", len(slice), cap(slice))
}
