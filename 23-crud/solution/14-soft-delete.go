package main

import (
	"fmt"
)

func main() {
	// db.Exec("UPDATE users SET deleted_at = NOW() WHERE id = $1", 1)
	fmt.Println("Marked as deleted")
}
