package main

import "fmt"

type Box struct {
	Value interface{}
}

func (b Box) ToString() string {
	return fmt.Sprintf("%v", b.Value)
}

func main() {
	b := Box{Value: 42}
	fmt.Printf("Converted: %s\n", b.ToString())
}
