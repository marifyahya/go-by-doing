package main

import (
	"fmt"
	"text/template"
	"bytes"
)

func main() {
	tpl := "Hello {{.Name}}!"
	t := template.Must(template.New("email").Parse(tpl))
	var buf bytes.Buffer
	t.Execute(&buf, map[string]string{"Name": "Alice"})
	fmt.Println("Template rendered")
}
