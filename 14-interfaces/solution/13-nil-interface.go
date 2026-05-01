package main

import "fmt"

type Shape interface {
	Area() int
}

func main() {
	var s Shape = nil
	if s == nil {
		fmt.Println("Nil interface")
	}
}
