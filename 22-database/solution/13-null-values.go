package main

import (
	"database/sql"
	"fmt"
)

func main() {
	var s sql.NullString
	s.Scan("value")
	if s.Valid {
		fmt.Println("Optional field handled")
	}
}
