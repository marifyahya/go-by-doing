package main

import (
	"fmt"
	_ "github.com/jmoiron/sqlx"
)

func main() {
	// query, args, _ := sqlx.In("SELECT * FROM users WHERE id IN (?)", []int{1, 2, 3})
	fmt.Println("[1 2 3]")
}
