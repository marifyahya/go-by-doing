package main

import (
	"errors"
	"fmt"
)

func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	res1, err1 := safeDivide(10, 2)
	fmt.Printf("Result: %d, Error: %v\n", res1, err1)

	res2, err2 := safeDivide(10, 0)
	fmt.Printf("Result: %d, Error: %v\n", res2, err2)
}
