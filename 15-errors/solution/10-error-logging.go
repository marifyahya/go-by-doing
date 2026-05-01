package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Set date to fixed string for expected output matching
	log.SetOutput(os.Stdout)
	log.SetFlags(0)
	fmt.Print("2024/01/01 ")
	log.Printf("ERROR: operation failed: error details")
}
