package main

import "fmt"

func generator() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func main() {
	for v := range generator() {
		fmt.Printf("%d ", v)
	}
	fmt.Println("...")
}
