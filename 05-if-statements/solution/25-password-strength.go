package main

import (
	"fmt"
	"unicode"
)

func main() {
	checkStrength := func(pw string) string {
		if len(pw) < 8 {
			return "Weak"
		}

		hasNumber := false
		hasUpper := false
		hasSpecial := false

		for _, char := range pw {
			if unicode.IsNumber(char) {
				hasNumber = true
			} else if unicode.IsUpper(char) {
				hasUpper = true
			} else if unicode.IsPunct(char) || unicode.IsSymbol(char) {
				hasSpecial = true
			}
		}

		if hasNumber && hasUpper && hasSpecial {
			return "Strong"
		} else if hasNumber {
			return "Medium"
		}
		return "Weak"
	}

	fmt.Printf("\"abc\": %s\n", checkStrength("abc"))
	fmt.Printf("\"abcd1234\": %s\n", checkStrength("abcd1234"))
	fmt.Printf("\"Abcd1234!\": %s\n", checkStrength("Abcd1234!"))
}
