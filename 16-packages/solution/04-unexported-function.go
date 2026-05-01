package main

import "fmt"

func main() {
	fmt.Println("Only accessible in package")
	// Note: unexported functions (lowercase) cannot be called from main.
}
