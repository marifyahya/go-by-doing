package main

import "fmt"

type Shape interface {
	Area() int
}

type Square struct {
	Side int
}

func (s Square) Area() int {
	return s.Side * s.Side
}

func main() {
	var s Shape = Square{Side: 5}
	fmt.Printf("Area: %d\n", s.Area())
}
