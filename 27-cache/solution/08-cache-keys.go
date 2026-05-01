package main

import (
	"fmt"
)

func main() {
	userID := 1
	key := fmt.Sprintf("user:%d:name", userID)
	fmt.Println(key)
}
