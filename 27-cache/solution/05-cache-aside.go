package main

import (
	"fmt"
)

func main() {
	// checkCache() -> if miss { fetchDB(); setCache() }
	fmt.Println("Cache miss, DB fetched")
}
