package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("nonexistent.txt")
	if os.IsNotExist(err) {
		fmt.Println("Not found")
	}
}
