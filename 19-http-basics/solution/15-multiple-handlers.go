package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {})
	fmt.Println("/ -> home")
	fmt.Println("/about -> about")
}
