package main

import "fmt"

func main() {
	fmt.Println("Started")
	defer fmt.Println("Cleanup done")
	fmt.Println("Finished")
}
