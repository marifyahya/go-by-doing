package main

import "fmt"

func main() {
	fmt.Println("Error: pointer arithmetic not allowed")
	// Go doesn't allow pointer arithmetic like p++ or p+1 directly.
}
