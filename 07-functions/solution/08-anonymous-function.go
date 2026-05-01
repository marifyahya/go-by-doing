package main

import "fmt"

func main() {
	func(x int) {
		fmt.Printf("Result: %d\n", x*x)
	}(5)
}
