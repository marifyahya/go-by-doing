package main

import "fmt"

func main() {
	n := 5
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	fmt.Printf("%d! = %d\n", n, fact)
}
