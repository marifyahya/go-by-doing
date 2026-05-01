package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}
