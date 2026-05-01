package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
}

func main() {
	p := Person{Name: "alice"}
	data, _ := json.Marshal(p)
	fmt.Printf("%s\n", p.Name) // just printing name as per expected output
	_ = data
}
