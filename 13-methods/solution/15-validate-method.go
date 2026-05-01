package main

import "fmt"

type User struct {
	Age int
}

func (u User) IsValid() bool {
	return u.Age >= 18
}

func main() {
	u1 := User{Age: 25}
	u2 := User{Age: 10}
	fmt.Printf("Valid: %t\n", u1.IsValid())
	fmt.Printf("Valid: %t\n", u2.IsValid())
}
