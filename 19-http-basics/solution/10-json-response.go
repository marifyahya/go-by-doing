package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	res := map[string]string{"message": "success"}
	data, _ := json.Marshal(res)
	fmt.Println(string(data))
}
