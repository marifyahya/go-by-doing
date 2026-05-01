package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := map[string]string{"name": "alice"}
	bytes, _ := json.Marshal(data)
	// rdb.Set(ctx, "user:1", bytes, 0)
	_ = bytes
	fmt.Println("JSON stored")
}
