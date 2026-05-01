package main

import "fmt"

func main() {
	m := map[string]int{"alice": 25}
	fmt.Println(m)
	delete(m, "alice")
	fmt.Printf("After delete: %v\n", m)
}
