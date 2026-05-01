package main

import (
	"fmt"
)

func main() {
	// db.Exec("INSERT INTO users(id, name) VALUES($1, $2) ON CONFLICT(id) DO UPDATE SET name = EXCLUDED.name", 1, "Alice")
	fmt.Println("Inserted/Updated")
}
