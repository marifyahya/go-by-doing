package main

import "fmt"

func main() {
	num := 12345
	res := 0
	for num > 0 {
		res = res*10 + num%10
		num /= 10
	}
	fmt.Printf("12345 reversed: %d\n", res)
}
