package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{"alice": 25, "bob": 30, "charlie": 35}
	var values []int
	for _, v := range m {
		values = append(values, v)
	}
	sort.Ints(values)
	fmt.Println(values)
}
