package main

import "fmt"

type Rectangle struct {
	Width, Height int
}

func (r Rectangle) Area() int {
	fmt.Println("[LOG] Area called")
	res := r.Width * r.Height
	fmt.Printf("Area: %d\n", res)
	fmt.Println("[LOG] Done")
	return res
}

func main() {
	r := Rectangle{10, 3}
	r.Area()
}
