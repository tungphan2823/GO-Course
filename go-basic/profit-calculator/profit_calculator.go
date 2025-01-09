package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter tax rate (in percentage): ")
	fmt.Scan(&taxRate)

	var earningB4Tax = revenue - expenses
	var earningAfterTax = earningB4Tax - (earningB4Tax * taxRate / 100)
	var ratio = earningB4Tax / earningAfterTax

	fmt.Printf("Earnings before tax: $%.2f\n", earningB4Tax)
	fmt.Printf("Earnings after tax: $%.2f\n", earningAfterTax)
	fmt.Printf("Rato: %.2f%%\n", ratio)
}
