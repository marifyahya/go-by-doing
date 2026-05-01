package main

import "fmt"

func main() {
	n := 10
	p := &n
	fmt.Printf("Address: %p\n", p)
	fmt.Printf("Value: %d\n", *p)
}
