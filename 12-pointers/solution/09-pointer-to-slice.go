package main

import "fmt"

func modifySlice(s []int) {
	s[2] = 99
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	modifySlice(slice)
	fmt.Println(slice)
}
