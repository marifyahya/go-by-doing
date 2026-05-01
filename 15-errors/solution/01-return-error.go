package main

import (
	"errors"
	"fmt"
)

func doSomething() error {
	return errors.New("something went wrong")
}

func main() {
	err := doSomething()
	fmt.Printf("Error: %v\n", err)
}
