package main

import "fmt"

func main() {
	m := map[string][]string{
		"fruits": {"apple", "orange"},
	}
	fmt.Printf("fruits: %v\n", m["fruits"])
}
