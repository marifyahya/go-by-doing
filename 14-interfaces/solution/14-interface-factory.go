package main

import "fmt"

type Shape interface {
	Area() int
}

type Square struct {
	Side int
}

func (s Square) Area() int { return s.Side * s.Side }

func NewShape(side int) Shape {
	return Square{Side: side}
}

func main() {
	s := NewShape(5)
	fmt.Printf("Area: %d\n", s.Area())
}
