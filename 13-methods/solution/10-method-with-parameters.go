package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(other Point) float64 {
	return math.Sqrt(math.Pow(p.X-other.X, 2) + math.Pow(p.Y-other.Y, 2))
}

func main() {
	p1 := Point{0, 0}
	p2 := Point{3, 4}
	fmt.Printf("Distance: %.0f\n", p1.Distance(p2))
}
