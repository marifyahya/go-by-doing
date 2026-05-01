package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{"alice": 25, "bob": 30}
	
	// Iteration order is random in Go. To match expected output reliably, 
	// one should ideally sort keys, but basic iteration is what's requested.
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("name: %s, age: %d\n", k, m[k])
	}
}
