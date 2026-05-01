package main

import (
	"fmt"
)

func logMessage(level string, msg string) {
	fmt.Printf("[%s] %s\n", level, msg)
}

func main() {
	logMessage("DEBUG", "debug message")
	logMessage("INFO", "info message")
	logMessage("ERROR", "error message")
}
