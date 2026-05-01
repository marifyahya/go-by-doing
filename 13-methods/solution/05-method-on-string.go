package main

import "fmt"

type MyString string

func (s MyString) Len() int {
	return len(s)
}

func main() {
	s := MyString("hello")
	fmt.Printf("%s: %d\n", s, s.Len())
}
