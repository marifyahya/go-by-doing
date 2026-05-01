package main

import "fmt"

func intSeq() func() int {
	i := 10
	return func() int {
		res := i
		i++
		return res
	}
}

func main() {
	next := intSeq()
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
}
