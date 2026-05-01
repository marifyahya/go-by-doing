package main

import "fmt"

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func main() {
	val, _ = divide(10, 2)
	fmt.Printf("Value: %d\n", val)
	fmt.Println("Ignored the rest")
}