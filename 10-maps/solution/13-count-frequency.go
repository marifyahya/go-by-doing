package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	input := []string{"a", "b", "a", "c", "b", "a"}
	freq := make(map[string]int)
	for _, s := range input {
		freq[s]++
	}

	keys := make([]string, 0, len(freq))
	for k := range freq {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var res []string
	for _, k := range keys {
		res = append(res, fmt.Sprintf("%s: %d", k, freq[k]))
	}
	fmt.Println(strings.Join(res, ", "))
}
