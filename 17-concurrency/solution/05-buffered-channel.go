package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	fmt.Printf("Channel with buffer %d\n", cap(ch))
}
