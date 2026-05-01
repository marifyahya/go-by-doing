package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type User struct {
	Name string
	Age  int
}

func main() {
	db, _ := sql.Open("postgres", "connStr")
	var u User
	// db.QueryRow(...).Scan(&u.Name, &u.Age)
	u = User{Name: "Alice", Age: 25}
	fmt.Printf("User: %s, Age: %d\n", u.Name, u.Age)
}
