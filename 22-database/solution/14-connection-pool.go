package main

import (
	"database/sql"
	"fmt"
)

func main() {
	db, _ := sql.Open("postgres", "connStr")
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	fmt.Println("Pool configured")
}
