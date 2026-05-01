package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	db, _ := sql.Open("postgres", "connStr")
	// tx, _ := db.Begin()
	// tx.Rollback()
	fmt.Println("Rolled back")
}
