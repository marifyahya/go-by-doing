package main

import (
	"fmt"
	"strings"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	var res []string
	for _, v := range arr {
		res = append(res, fmt.Sprintf("%d", v))
	}
	fmt.Printf("[%s]\n", strings.Join(res, ", "))
}
