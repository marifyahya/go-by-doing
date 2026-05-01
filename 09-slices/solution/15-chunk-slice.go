package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	chunkSize := 2
	var chunks [][]int
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	fmt.Println(chunks)
}
