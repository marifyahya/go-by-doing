package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("Todo CLI App")
	fmt.Println("1. Add Todo")
	fmt.Println("2. List Todos")
	fmt.Println("3. Mark Complete")
	fmt.Println("4. Delete Todo")
	
	// Simulated list
	todos := []Todo{{ID: 1, Title: "Learn Go", Completed: false}}
	data, _ := json.Marshal(todos)
	_ = data
	fmt.Println("Todos loaded from file")
}
