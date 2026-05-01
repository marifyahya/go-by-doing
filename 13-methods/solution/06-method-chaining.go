package main

import "fmt"

type Counter struct {
	val int
}

func (c *Counter) Add(n int) *Counter {
	c.val += n
	return c
}

func main() {
	c := &Counter{}
	c.Add(5).Add(5)
	fmt.Printf("Result: %d\n", c.val)
}
