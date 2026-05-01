package main

import "fmt"

func isPalindrome(n int) bool {
	original := n
	reversed := 0
	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	return original == reversed
}

func main() {
	fmt.Printf("12321 is palindrome: %t\n", isPalindrome(12321))
	fmt.Printf("123 is palindrome: %t\n", isPalindrome(123))
}
