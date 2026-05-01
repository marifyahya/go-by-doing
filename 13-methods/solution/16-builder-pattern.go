package main

import "fmt"

type User struct {
	Name  string
	Age   int
	Admin bool
}

type UserBuilder struct {
	u User
}

func (b *UserBuilder) SetName(name string) *UserBuilder {
	b.u.Name = name
	return b
}

func (b *UserBuilder) SetAge(age int) *UserBuilder {
	b.u.Age = age
	return b
}

func (b *UserBuilder) SetAdmin(admin bool) *UserBuilder {
	b.u.Admin = admin
	return b
}

func (b *UserBuilder) Build() User {
	return b.u
}

func main() {
	u := (&UserBuilder{}).SetName("Alice").SetAge(25).SetAdmin(true).Build()
	fmt.Printf("User: %s, Age: %d, Admin: %t\n", u.Name, u.Age, u.Admin)
}
