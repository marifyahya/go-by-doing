package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	original := map[string]int{"a": 1, "b": 2, "c": 3}
	filtered := make(map[string]int)
	for k, v := range original {
		if v > 1 {
			filtered[k] = v
		}
	}

	fmt.Printf("Original: %v\n", original)
	
	keys := make([]string, 0, len(filtered))
	for k := range filtered {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	
	var res []string
	for _, k := range keys {
		res = append(res, fmt.Sprintf("%s:%d", k, filtered[k]))
	}
	fmt.Printf("Filtered (>1): {%s}\n", strings.Join(res, ", "))
}
