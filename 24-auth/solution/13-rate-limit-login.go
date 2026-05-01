package main

import (
	"fmt"
)

func main() {
	attempts := 6
	if attempts > 5 {
		fmt.Println("Too many attempts")
	}
}
