package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func find() error {
	return ErrNotFound
}

func main() {
	err := find()
	if errors.Is(err, ErrNotFound) {
		fmt.Println("ErrNotFound returned")
	}
}
