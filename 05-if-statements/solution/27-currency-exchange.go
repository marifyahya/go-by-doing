package main

import "fmt"

func main() {
	amount := 100.0
	currency := "EUR"
	var result float64

	if currency == "EUR" {
		result = amount * 0.92
	} else if currency == "JPY" {
		result = amount * 149
	}

	fmt.Printf("%.0f USD to %s: %.0f\n", amount, currency, result)
}
