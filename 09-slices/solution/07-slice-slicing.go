package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	sub := slice[1:4]
	fmt.Printf("Original: %v\n", slice)
	fmt.Printf("Sub: %v\n", sub)
}
