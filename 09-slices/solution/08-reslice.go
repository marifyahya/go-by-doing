package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	reslice := slice[0:3]
	fmt.Println(reslice)
	reslice = reslice[0:5] // possible because underlying array has 5 elements
	fmt.Println(reslice)
}
