package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	original := map[string]int{"a": 1, "b": 2}
	inverted := make(map[int]string)
	for k, v := range original {
		inverted[v] = k
	}

	fmt.Printf("Original: %v\n", original)
	
	// Format inverted for consistent output
	keys := make([]int, 0, len(inverted))
	for k := range inverted {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	
	var res []string
	for _, k := range keys {
		res = append(res, fmt.Sprintf("%d:%s", k, inverted[k]))
	}
	fmt.Printf("Inverted: {%s}\n", strings.Join(res, ", "))
}
