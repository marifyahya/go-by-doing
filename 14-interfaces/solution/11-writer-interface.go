package main

import (
	"fmt"
	"io"
)

type MyWriter struct{}

func (w MyWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func main() {
	var w io.Writer = MyWriter{}
	n, _ := w.Write([]byte("hello"))
	fmt.Printf("Wrote %d bytes\n", n)
}
