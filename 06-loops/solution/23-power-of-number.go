package main

import "fmt"

func main() {
	base := 2
	exp := 10
	res := 1
	for i := 0; i < exp; i++ {
		res *= base
	}
	fmt.Printf("%d^%d = %d\n", base, exp, res)
}
