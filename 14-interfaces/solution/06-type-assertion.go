package main

import "fmt"

func main() {
	var i interface{} = 42
	val, ok := i.(int)
	if ok {
		fmt.Printf("Type is int, Value: %d\n", val)
	}
}
