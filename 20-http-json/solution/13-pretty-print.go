package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	u := map[string]string{"name": "Alice"}
	data, _ := json.MarshalIndent(u, "", "  ")
	fmt.Println(string(data))
}
