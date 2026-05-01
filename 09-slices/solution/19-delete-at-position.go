package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	pos := 2
	res := append(slice[:pos], slice[pos+1:]...)
	fmt.Printf("After delete at %d: %v\n", pos, res)
}
