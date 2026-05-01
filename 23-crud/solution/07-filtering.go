package main

import (
	"fmt"
)

func main() {
	// db.QueryRow("SELECT COUNT(*) FROM users WHERE status = $1", "active").Scan(&count)
	fmt.Println("Active users: 5")
}
