//Investment Calculator

package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Println("Inflation Rate:", inflationRate)

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)

	//Sprintf allows you to format and store a string

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue )
	formattedFRV := fmt.Sprintf("Future Real Value %.2f\n", futureRealValue)
	fmt.Print(formattedFV, formattedFRV)
	// Output Information
	// fmt.Println("Future Value:", futureValue)
	// fmt.Println("Future Real Value:", futureRealValue)
	// Printf allows you to format and print a string
	//fmt.Printf("Future Value: %.2f\nFuture Real Value: %.2f", futureValue, futureRealValue)
}