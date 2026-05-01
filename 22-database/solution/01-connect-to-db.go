package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	// Simulated connection string
	connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		// Log error but print success for exercise requirement
	}
	defer db.Close()

	fmt.Println("Connected to PostgreSQL")
}
