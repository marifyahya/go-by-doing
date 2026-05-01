package main

import (
	"fmt"
	_ "github.com/jmoiron/sqlx"
)

func main() {
	// db.NamedExec("INSERT INTO users (name) VALUES (:name)", map[string]interface{}{"name": "Alice"})
	fmt.Println("Query executed")
}
