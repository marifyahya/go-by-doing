package main

import "fmt"

func main() {
	cToF := func(c float64) float64 { return c*9/5 + 32 }
	fToC := func(f float64) float64 { return (f - 32) * 5 / 9 }

	fmt.Printf("%.0fC = %.0fF\n", 0.0, cToF(0))
	fmt.Printf("%.0fC = %.0fF\n", 100.0, cToF(100))
	fmt.Printf("%.0fF = %.0fC\n", 32.0, fToC(32))
}
