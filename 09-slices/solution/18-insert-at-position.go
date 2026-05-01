package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	pos := 2
	val := 99
	res := append(slice[:pos], append([]int{val}, slice[pos:]...)...)
	fmt.Printf("After insert %d at %d: %v\n", val, pos, res)
}
