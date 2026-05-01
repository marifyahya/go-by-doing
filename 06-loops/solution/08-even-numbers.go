package main

import "fmt"

func main() {
	fmt.Print("Even numbers up to 10: ")
	for i := 2; i <= 10; i += 2 {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}
