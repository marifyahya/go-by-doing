package main

import (
	"fmt"
	"net/http"
)

func main() {
	_ = http.FileServer(http.Dir("./static"))
	fmt.Println("CSS/JS served")
}
