package main

import "fmt"

type Shape interface {
	Area() int
	Perimeter() int
}

type Square struct {
	Side int
}

func (s Square) Area() int      { return s.Side * s.Side }
func (s Square) Perimeter() int { return 4 * s.Side }

func main() {
	var s Shape = Square{Side: 5}
	fmt.Printf("Area: %d, Perimeter: %d\n", s.Area(), s.Perimeter())
}
