package main

import "fmt"

func main() {
	var a int64 = 1000000
	var b int64 = 1000000
	result := a * b

	fmt.Printf("%d * %d = %d\n", a, b, result)
}