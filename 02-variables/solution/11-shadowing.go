package main

import "fmt"

func main() {
	x := 1
	fmt.Printf("Outer: %d\n", x)

	{
		x := 2
		fmt.Printf("Inner: %d\n", x)
	}

	fmt.Printf("Outer after: %d\n", x)
}