package main

import (
	"fmt"
)

func main() {
	page, limit := 1, 3
	offset := (page - 1) * limit
	// db.Query("SELECT name FROM users LIMIT $1 OFFSET $2", limit, offset)
	_ = offset
	fmt.Println("Page 1 of 3")
	fmt.Println("[user1, user2, user3]")
}
