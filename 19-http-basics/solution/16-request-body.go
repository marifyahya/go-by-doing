package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	body := strings.NewReader(strings.Repeat("a", 100))
	content, _ := io.ReadAll(body)
	fmt.Printf("Body length: %d\n", len(content))
}
