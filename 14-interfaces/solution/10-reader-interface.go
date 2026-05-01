package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("hello")
	p := make([]byte, 5)
	n, _ := r.Read(p)
	fmt.Printf("Read %d bytes\n", n)
}
