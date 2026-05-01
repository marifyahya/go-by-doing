package main

import (
	"fmt"
	"sync"
)

func main() {
	var pool = &sync.Pool{
		New: func() interface{} { return "resource" },
	}
	res := pool.Get()
	fmt.Printf("Pooled %s used\n", res)
	pool.Put(res)
}
