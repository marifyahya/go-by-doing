package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Content-Type: application/json")
	res := map[string]string{"data": "value"}
	json.NewEncoder(os.Stdout).Encode(res)
}
