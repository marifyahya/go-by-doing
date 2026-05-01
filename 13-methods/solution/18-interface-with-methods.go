package main

import "fmt"

type Shape interface {
	Area() int
}

type Rectangle struct {
	Width, Height int
}

func (r Rectangle) Area() int {
	return r.Width * r.Height
}

func main() {
	var s Shape = Rectangle{10, 3}
	fmt.Printf("Area: %d\n", s.Area())
}
