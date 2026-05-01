package main

import (
	"errors"
	"fmt"
)

func main() {
	errFile := errors.New("file")
	errRead := fmt.Errorf("read -> %w", errFile)
	errProcess := fmt.Errorf("process -> %w", errRead)

	fmt.Printf("All errors: %v\n", errProcess)
}
