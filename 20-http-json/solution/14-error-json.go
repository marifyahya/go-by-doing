package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	res := map[string]string{"error": "message"}
	data, _ := json.Marshal(res)
	fmt.Println(string(data))
}
