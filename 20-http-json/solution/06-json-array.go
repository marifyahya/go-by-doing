package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	data, _ := json.Marshal(arr)
	fmt.Println(string(data))
}
