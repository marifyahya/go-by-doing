package main

import "fmt"

func main() {
	a, b, c := 1, 2, 3
	ptrs := [3]*int{&a, &b, &c}
	for _, p := range ptrs {
		fmt.Printf("%d ", *p)
	}
	fmt.Println()
}
