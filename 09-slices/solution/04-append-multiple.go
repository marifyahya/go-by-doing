package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	vals := []int{4, 5, 6}
	slice = append(slice, vals...)
	fmt.Println(slice)
}
