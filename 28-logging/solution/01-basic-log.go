package main

import (
	"log"
	"os"
)

func main() {
	l := log.New(os.Stdout, "", 0)
	l.Println("2024/01/01 12:00:00 INFO Starting application")
}
