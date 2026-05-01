package main

import "fmt"

func main() {
	fmt.Println("Countdown: 5")
	i := 4
	for i >= 0 {
		fmt.Println(i)
		i--
	}
}
