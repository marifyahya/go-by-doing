package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := []byte(`{"name":"Alice"}`)
	fmt.Printf("Valid: %t\n", json.Valid(data))
}
