package main

import "fmt"

type MyError struct{}

func (e MyError) Error() string {
	return "custom error"
}

func main() {
	var err error = MyError{}
	fmt.Printf("Error: %v\n", err)
}
