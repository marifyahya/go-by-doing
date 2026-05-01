package main

import "fmt"

type IntSlice []int

func (s IntSlice) Sum() int {
	total := 0
	for _, v := range s {
		total += v
	}
	return total
}

func main() {
	s := IntSlice{1, 2, 3, 4, 5}
	fmt.Printf("Sum: %d\n", s.Sum())
}
