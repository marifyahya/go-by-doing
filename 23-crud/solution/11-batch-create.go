package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		// db.Exec("INSERT INTO users(name) VALUES($1)", fmt.Sprintf("user%d", i))
	}
	fmt.Println("10 users created")
}
