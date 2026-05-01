package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	UserID int `json:"userID"`
	jwt.RegisteredClaims
}

func main() {
	claims := MyClaims{UserID: 123}
	fmt.Printf("UserID: %d\n", claims.UserID)
}
