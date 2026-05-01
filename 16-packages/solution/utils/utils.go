package utils

import "fmt"

func init() {
	fmt.Println("Utils loaded")
}

func SayHello() {
	fmt.Println("Hello from utils")
}

func Exported() {
	fmt.Println("Exported function called")
}

func internal() {
	fmt.Println("Only accessible in package")
}
