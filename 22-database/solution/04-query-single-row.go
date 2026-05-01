package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	db, _ := sql.Open("postgres", "connStr")
	var name string
	// err := db.QueryRow("SELECT name FROM users WHERE id = $1", 1).Scan(&name)
	name = "Alice"
	fmt.Printf("Name: %s\n", name)
}
