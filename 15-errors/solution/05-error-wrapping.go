package main

import (
	"errors"
	"fmt"
)

func main() {
	original := errors.New("original error")
	wrapped := fmt.Errorf("process failed: %w", original)
	fmt.Println(wrapped)
}
