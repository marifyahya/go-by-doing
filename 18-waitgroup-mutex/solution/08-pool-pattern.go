package main

import (
	"fmt"
	"sync"
)

func main() {
	p := &sync.Pool{
		New: func() interface{} { return "Pooled resource used" },
	}
	fmt.Println(p.Get())
}
