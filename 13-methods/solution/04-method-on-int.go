package main

import "fmt"

type MyInt int

func (i MyInt) IsEven() bool {
	return i%2 == 0
}

func main() {
	val := MyInt(5)
	fmt.Printf("%d is even: %t\n", val, val.IsEven())
}
