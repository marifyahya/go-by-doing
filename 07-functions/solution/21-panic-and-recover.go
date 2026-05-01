package main

import "fmt"

func mayPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()
	panic("something went wrong")
}

func main() {
	mayPanic()
}
