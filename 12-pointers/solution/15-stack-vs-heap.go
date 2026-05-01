package main

import "fmt"

func main() {
	fmt.Println("Value on heap")
	// Returning a pointer to a local variable causes it to be allocated on the heap (escape analysis).
}
