package main

import "fmt"

type Rectangle struct {
	Width, Height int
}

func (r Rectangle) Area() int      { return r.Width * r.Height }
func (r Rectangle) Perimeter() int { return 2 * (r.Width + r.Height) }

func main() {
	r := Rectangle{Width: 10, Height: 3}
	fmt.Printf("Area: %d\n", r.Area())
	fmt.Printf("Perimeter: %d\n", r.Perimeter())
}
