package main

import "fmt"

type ValidationError struct {
	Message string
}

func (e ValidationError) Error() string {
	return "ValidationError: " + e.Message
}

func main() {
	err := ValidationError{Message: "name is required"}
	fmt.Println(err.Error())
}
