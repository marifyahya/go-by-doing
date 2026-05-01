package main

import "fmt"

func main() {
	square := func(x int) int {
		return x * x
	}
	fmt.Printf("Square: %d\n", square(4))
}
