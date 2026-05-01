package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	n := 2
	rotated := append(slice[n:], slice[:n]...)
	fmt.Printf("After rotate %d: %v\n", n, rotated)
}
