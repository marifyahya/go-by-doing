package main

import (
	"fmt"
	"sort"
)

type MyMap map[string]int

func (m MyMap) Keys() []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func main() {
	m := MyMap{"a": 1, "b": 2, "c": 3}
	fmt.Printf("Keys: %v\n", m.Keys())
}
