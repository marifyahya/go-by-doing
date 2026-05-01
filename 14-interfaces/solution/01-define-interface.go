package main

import "fmt"

type Shape interface {
	Area() int
}

func main() {
	fmt.Println("Shape has Area method")
}
