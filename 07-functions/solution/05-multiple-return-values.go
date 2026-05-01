package main

import "fmt"

func divide(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	q, r := divide(5, 2)
	fmt.Printf("Quotient: %d, Remainder: %d\n", q, r)
}
