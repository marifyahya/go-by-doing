package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	const jsonStream = `
		{"Name": "Alice", "Age": 25}
		{"Name": "Bob", "Age": 30}
	`
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for dec.More() {
		var m map[string]interface{}
		err := dec.Decode(&m)
		if err != nil {
			break
		}
	}
	fmt.Println("Stream processed")
}
