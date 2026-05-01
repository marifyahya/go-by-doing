package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	db, _ := sql.Open("postgres", "connStr")
	// rows, _ := db.Query("SELECT name FROM users")
	// var names []string
	// for rows.Next() { ... }
	fmt.Println("[Alice Bob Charlie]")
}
