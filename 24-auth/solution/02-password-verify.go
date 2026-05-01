package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "secret123"
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	fmt.Printf("Valid: %t\n", err == nil)
}
