package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
}

func main() {
	data := []byte(`{"name":"Alice"}`)
	var p Person
	json.Unmarshal(data, &p)
	fmt.Println(p.Name)
}
