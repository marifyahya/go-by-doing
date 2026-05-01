package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	m1 := map[string]int{"first": 1, "second": 2}
	m2 := map[string]int{"third": 3}
	
	merged := make(map[string]int)
	for k, v := range m1 {
		merged[k] = v
	}
	for k, v := range m2 {
		merged[k] = v
	}

	keys := make([]string, 0, len(merged))
	for k := range merged {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	
	var res []string
	for _, k := range keys {
		res = append(res, fmt.Sprintf("%s: %d", k, merged[k]))
	}
	fmt.Printf("{%s}\n", strings.Join(res, ", "))
}
