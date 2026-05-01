package main

import "fmt"

type ValidationError struct{}

func (e ValidationError) Error() string { return "Validation error" }

func main() {
	var err error = ValidationError{}
	switch err.(type) {
	case ValidationError:
		fmt.Println("Validation error")
	default:
		fmt.Println("Unknown error")
	}
}
