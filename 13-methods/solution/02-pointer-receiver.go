package main

import "fmt"

type Rectangle struct {
	Width, Height int
}

func (r Rectangle) Area() int {
	return r.Width * r.Height
}

func (r *Rectangle) ScaleHeight(f int) {
	r.Height *= f
}

func main() {
	r := &Rectangle{Width: 10, Height: 3}
	fmt.Printf("Area: %d\n", r.Area())
	r.ScaleHeight(2) // 10 * 6 = 60
	fmt.Printf("After: %d\n", r.Area())
}
