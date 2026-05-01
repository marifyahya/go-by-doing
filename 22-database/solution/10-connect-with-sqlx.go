package main

import (
	"fmt"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	// db, _ := sqlx.Connect("postgres", "connStr")
	fmt.Println("sqlx connected")
}
