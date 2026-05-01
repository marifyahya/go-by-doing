package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{"alice": 25, "bob": 30, "charlie": 35}
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Println(keys)
}
