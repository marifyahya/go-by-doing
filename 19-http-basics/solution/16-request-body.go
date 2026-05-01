package main

import (
	"fmt"
	"io"
	"net/http/httptest"
	"strings"
)

func main() {
	body := strings.Repeat("a", 100)
	req := httptest.NewRequest("POST", "/", strings.NewReader(body))
	b, _ := io.ReadAll(req.Body)
	fmt.Printf("Body length: %d\n", len(b))
}
