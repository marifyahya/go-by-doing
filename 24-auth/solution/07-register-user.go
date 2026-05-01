package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	username := "alice"
	password := "secret"
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// saveToDB(username, hash)
	_ = hash
	fmt.Printf("User registered: %s\n", username)
}
