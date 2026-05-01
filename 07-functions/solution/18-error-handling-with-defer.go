package main

import (
	"errors"
	"fmt"
)

func process() error {
	defer fmt.Println("Cleanup done")
	return errors.New("file not found")
}

func main() {
	err := process()
	if err != nil {
		fmt.Printf("Error occurred: %v\n", err)
	}
}
