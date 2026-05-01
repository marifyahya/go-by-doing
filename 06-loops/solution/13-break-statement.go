package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 4 {
			break
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}
