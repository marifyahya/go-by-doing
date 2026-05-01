package main

import "fmt"

func main() {
	var i interface{} = 42
	switch v := i.(type) {
	case int:
		fmt.Println("Type: int")
	case string:
		fmt.Println("Type: string")
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}
