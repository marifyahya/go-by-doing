package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello world"
	vowels := "aeiouAEIOU"
	count := 0
	for _, char := range str {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	fmt.Printf("\"%s\" has %d vowels\n", str, count)
}
