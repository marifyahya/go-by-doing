package main

import (
	"database/sql"
	"fmt"
)

func main() {
	var name, email string
	// db.QueryRow("SELECT name, email FROM users WHERE id = $1", 1).Scan(&name, &email)
	name, email = "Alice", "alice@test.com"
	fmt.Printf("User: %s, Email: %s\n", name, email)
}
