package main

import (
	"fmt"
)

func validate(email string) error {
	if email == "" {
		return fmt.Errorf("email required")
	}
	return nil
}

func main() {
	err := validate("")
	if err != nil {
		fmt.Printf("Validation failed: %v\n", err)
	}
}
