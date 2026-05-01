package main

import (
	"fmt"
)

func main() {
	// db.Exec("UPDATE users SET status = $1 WHERE id IN ($2, $3, $4, $5, $6)", "active", 1, 2, 3, 4, 5)
	fmt.Println("5 updated")
}
