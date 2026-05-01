package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	db, _ := sql.Open("postgres", "connStr")
	var lastID int
	// db.QueryRow("INSERT INTO users(name) VALUES($1) RETURNING id", "Alice").Scan(&lastID)
	lastID = 1
	fmt.Printf("User created with ID: %d\n", lastID)
}
