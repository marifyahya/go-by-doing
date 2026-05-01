package main

import "fmt"

type Areaer interface {
	Area() int
}

type Perimeterer interface {
	Perimeter() int
}

type Shape interface {
	Areaer
	Perimeterer
}

func main() {
	fmt.Println("Has Area and Perimeter")
}
