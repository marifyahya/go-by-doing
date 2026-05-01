package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	valid := json.Valid([]byte(`{"name":"Alice"}`))
	fmt.Printf("Valid: %t\n", valid)
}
