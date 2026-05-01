package main

import "fmt"

func main() {
	i := 42
	f := 3.14
	s := "42"

	iToF := float64(i)
	fToI := int(f)
	sToI := string(s[0])

	fmt.Printf("int to float: %.2f\n", iToF)
	fmt.Printf("float to int: %d\n", fToI)
	fmt.Printf("int to string: %s\n", string(sToI))
}