package main

import (
	"errors"
	"fmt"
)

type User struct {
	Name string
}

func (u User) Validate() error {
	if u.Name == "" {
		return errors.New("name required")
	}
	return nil
}

func main() {
	u1 := User{Name: "Alice"}
	if err := u1.Validate(); err == nil {
		fmt.Println("Valid")
	}

	u2 := User{Name: ""}
	if err := u2.Validate(); err != nil {
		fmt.Printf("Invalid: %v\n", err)
	}
}
