package main

import "fmt"

func main() {
	num := 123
	sum := 0
	n := num
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	fmt.Printf("%d sum of digits: %d\n", num, sum)
}
