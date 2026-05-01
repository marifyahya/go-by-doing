package main

import "fmt"

func main() {
	n := 10
	p := &n
	pp := &p
	fmt.Printf("Value: %d\n", **pp)
}
