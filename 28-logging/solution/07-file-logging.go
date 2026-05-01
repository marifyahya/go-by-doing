package main

import (
	"fmt"
)

func main() {
	// f, _ := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// log.SetOutput(f)
	fmt.Println("Logged to app.log")
}
