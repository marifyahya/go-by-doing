package main

import "fmt"

func main() {
	var i interface{} = 42
	fmt.Printf("Value: %v\n", i)
	i = "hello"
	fmt.Printf("Value: %v\n", i)
}
