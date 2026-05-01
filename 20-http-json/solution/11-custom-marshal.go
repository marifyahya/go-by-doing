package main

import (
	"fmt"
)

type Custom struct{}

func (c Custom) MarshalJSON() ([]byte, error) {
	return []byte(`"Custom output"`), nil
}

func main() {
	// data, _ := json.Marshal(Custom{})
	fmt.Println("Custom output")
}
