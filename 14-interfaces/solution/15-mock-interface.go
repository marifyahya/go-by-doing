package main

import "fmt"

type Service interface {
	Do()
}

type MockService struct{}

func (m MockService) Do() {
	fmt.Println("Mock called")
}

func main() {
	var s Service = MockService{}
	s.Do()
}
