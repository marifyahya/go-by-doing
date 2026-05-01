package main

import "fmt"

func main() {
	x := 10
	fmt.Printf("x = %d\n", x)

	x += 5
	fmt.Printf("x += 5: %d\n", x)

	x -= 3
	fmt.Printf("x -= 3: %d\n", x)

	x *= 2
	fmt.Printf("x *= 2: %d\n", x)
}