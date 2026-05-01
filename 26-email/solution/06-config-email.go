package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("SMTP_HOST", "smtp.gmail.com")
	fmt.Println("Configured")
}
