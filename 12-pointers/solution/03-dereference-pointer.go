package main

import "fmt"

func main() {
	n := 10
	p := &n
	fmt.Printf("Value: %d\n", *p)
}
