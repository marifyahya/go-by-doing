package main

import (
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "APP: ", 0)
	logger.Println("Message")
}
