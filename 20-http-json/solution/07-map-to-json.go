package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := map[string]string{"key1": "value1", "key2": "value2"}
	data, _ := json.Marshal(m)
	fmt.Println(string(data))
}
