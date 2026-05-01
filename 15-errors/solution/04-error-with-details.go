package main

import "fmt"

type InputError struct {
	Field  string
	Reason string
}

func (e InputError) Error() string {
	return fmt.Sprintf("InvalidInput: field=%s, reason=%s", e.Field, e.Reason)
}

func main() {
	err := InputError{Field: "name", Reason: "required"}
	fmt.Println(err.Error())
}
