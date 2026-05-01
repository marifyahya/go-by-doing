package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	log := map[string]string{
		"level":   "info",
		"message": "request",
		"ip":      "127.0.0.1",
	}
	data, _ := json.Marshal(log)
	fmt.Println(string(data))
}
