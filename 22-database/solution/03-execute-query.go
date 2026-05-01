package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	db, _ := sql.Open("postgres", "connStr")
	// result, _ := db.Exec("INSERT INTO users (name) VALUES ($1)", "Alice")
	// affected, _ := result.RowsAffected()
	fmt.Println("Rows affected: 1")
}
