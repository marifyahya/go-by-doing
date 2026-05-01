package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	db, _ := sql.Open("postgres", "connStr")
	// tx, _ := db.Begin()
	// tx.Exec(...)
	// tx.Commit()
	fmt.Println("Transaction committed")
}
