package main

import "fmt"

func main() {
	slice := make([]int, 3, 3)
	fmt.Printf("Length: %d, Capacity: %d\n", len(slice), cap(slice))
	fmt.Println(slice)
}
