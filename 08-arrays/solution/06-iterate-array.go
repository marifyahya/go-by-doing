package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	for i, v := range arr {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
}
