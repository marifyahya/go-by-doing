package main

import "fmt"

func main() {
	n := 10
	p := &n
	fmt.Printf("Before: %d\n", n)
	*p = 20
	fmt.Printf("After: %d\n", n)
}
