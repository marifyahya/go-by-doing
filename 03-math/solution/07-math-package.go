package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Absolute: %d\n", int(math.Abs(-5)))
	fmt.Printf("Square root: %.0f\n", math.Sqrt(9))
	fmt.Printf("Power: %.0f\n", math.Pow(2, 3))
	fmt.Printf("Ceil: %.0f\n", math.Ceil(3.14))
	fmt.Printf("Floor: %.0f\n", math.Floor(3.14))
}