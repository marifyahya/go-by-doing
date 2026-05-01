package main

import (
	"errors"
	"fmt"
)

func doSomething() error {
	return errors.New("something went wrong")
}

func main() {
	if err := doSomething(); err != nil {
		fmt.Printf("Failed: %v\n", err)
	}
}
