package main

import "fmt"

func main() {
	checkTriangle := func(a, b, c int) string {
		if a+b <= c || a+c <= b || b+c <= a {
			return "Invalid"
		}
		if a == b && b == c {
			return "Equilateral"
		}
		if a == b || b == c || a == c {
			return "Isosceles"
		}
		return "Scalene"
	}

	fmt.Printf("3, 3, 3: %s\n", checkTriangle(3, 3, 3))
	fmt.Printf("3, 4, 5: %s\n", checkTriangle(3, 4, 5))
	fmt.Printf("3, 3, 5: %s\n", checkTriangle(3, 3, 5))
}
