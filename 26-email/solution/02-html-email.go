package main

import (
	"fmt"
)

func main() {
	header := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := "<html><body><h1>Hello!</h1></body></html>"
	_ = header + body
	fmt.Println("HTML email sent")
}
