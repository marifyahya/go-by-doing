package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	db, _ := sql.Open("postgres", "connStr")
	// stmt, _ := db.Prepare("INSERT INTO users(name) VALUES($1)")
	// stmt.Exec("Alice")
	fmt.Println("Prepared executed")
}
