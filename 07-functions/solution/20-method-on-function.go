package main

import "fmt"

type CalcFunc func(int) int

func (f CalcFunc) Execute(n int) int {
	return f(n)
}

func main() {
	var double CalcFunc = func(n int) int {
		return n * 2
	}
	fmt.Printf("Calculation: %d\n", double.Execute(10))
}
