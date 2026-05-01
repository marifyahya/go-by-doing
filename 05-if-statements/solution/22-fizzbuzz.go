package main

import (
	"fmt"
	"strings"
)

func main() {
	var results []string
	for i := 1; i <= 15; i++ {
		if i%3 == 0 && i%5 == 0 {
			results = append(results, "FizzBuzz")
		} else if i%3 == 0 {
			results = append(results, "Fizz")
		} else if i%5 == 0 {
			results = append(results, "Buzz")
		} else {
			results = append(results, fmt.Sprintf("%d", i))
		}
	}
	fmt.Println(strings.Join(results, ", "))
}
