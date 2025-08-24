// Profit Calculator
package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	earnings_before_tax := revenue - expenses
	profit := earnings_before_tax * (1 - taxRate / 100)
	ratio := earnings_before_tax / profit

	fmt.Println("Earnings Before Tax:", earnings_before_tax)
	fmt.Println("Profit:", profit)
	fmt.Print("Ratio: ", ratio)
}