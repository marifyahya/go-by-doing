package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Msg string
}

func (e ValidationError) Error() string { return e.Msg }

func main() {
	err := ValidationError{Msg: "name required"}
	var vErr ValidationError
	if errors.As(err, &vErr) {
		fmt.Printf("Validation error: %v\n", vErr.Msg)
	}
}
