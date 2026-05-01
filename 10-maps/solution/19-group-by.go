package main

import (
	"fmt"
	"sort"
)

func main() {
	people := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 10,
	}

	grouped := make(map[string][]string)
	for name, age := range people {
		category := "Adult"
		if age < 18 {
			category = "Child"
		}
		grouped[category] = append(grouped[category], name)
	}

	// Sort for consistency
	categories := []string{"Adult", "Child"}
	for _, cat := range categories {
		names := grouped[cat]
		sort.Strings(names)
		fmt.Printf("%s: %v\n", cat, names)
	}
}
