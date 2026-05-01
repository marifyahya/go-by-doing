package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["alice"] = 25
	fmt.Println(m)
	m["bob"] = 30
	fmt.Println(m)
}
