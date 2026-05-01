package main

import "fmt"

func main() {
	arr := [5]int{10, 20, 30, 40, 50}
	fmt.Printf("First: %d\n", arr[0])
	fmt.Printf("Last: %d\n", arr[len(arr)-1])
}
