package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

func main() {
	raw := json.RawMessage(`{"key":"value"}`)
	m := Message{Type: "test", Data: raw}
	_ = m
	fmt.Println("Raw JSON preserved")
}
