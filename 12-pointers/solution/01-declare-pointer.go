package main

import "fmt"

func main() {
	var n int = 10
	var p *int = &n
	fmt.Printf("%p\n", p)
}
