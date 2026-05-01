package main

import (
	"fmt"
)

func main() {
	err := fmt.Errorf("connection refused")
	fmt.Printf("ERROR: failed to connect: %v\n", err)
}
