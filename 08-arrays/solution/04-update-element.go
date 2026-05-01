package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	arr[2] = 99
	fmt.Println(arr)
}
