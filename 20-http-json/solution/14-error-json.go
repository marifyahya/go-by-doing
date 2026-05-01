package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	errLog := map[string]string{"error": "message"}
	data, _ := json.Marshal(errLog)
	fmt.Println(string(data))
}
