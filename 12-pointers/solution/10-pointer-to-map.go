package main

import "fmt"

func modifyMap(m map[string]int) {
	m["a"] = 1
	m["b"] = 2
}

func main() {
	m := make(map[string]int)
	modifyMap(m)
	fmt.Println(m)
}
