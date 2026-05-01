package main

import (
	"fmt"
	"strings"
)

func generator() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	next := generator()
	var res []string
	for i := 0; i < 5; i++ {
		res = append(res, fmt.Sprintf("%d", next()))
	}
	fmt.Println(strings.Join(res, ", "))
}
