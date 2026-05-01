package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func main() {
	db, _ := sql.Open("postgres", "connStr")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	
	// db.QueryContext(ctx, "SELECT * FROM users")
	_ = ctx
	fmt.Println("Query with timeout")
}
