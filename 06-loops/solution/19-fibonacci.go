package main

import (
	"fmt"
	"strings"
)

func main() {
	n := 10
	fib := make([]int, n)
	fib[0], fib[1] = 0, 1
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	var results []string
	for _, v := range fib {
		results = append(results, fmt.Sprintf("%d", v))
	}
	fmt.Println(strings.Join(results, ", "))
}
