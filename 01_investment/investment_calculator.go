//Investment Calculator

package main

import (
	"fmt"
	"math"
)

// global variable. moved const inflationRate = 2.5 out of main function because it will stay the same
const inflationRate = 2.5

func main() {
// scoped variables left in the main function because they have to be set up when main function is sxecuted
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	// fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Println("Inflation Rate:", inflationRate)

	// futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	// futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)
	//here we can function and store two returned values in two variables
	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, inflationRate )

	//Sprintf allows you to format and store a string

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue )
	formattedFRV := fmt.Sprintf("Future Real Value %.2f", futureRealValue)

	fmt.Print(formattedFV, formattedFRV)
	// Output Information
	// fmt.Println("Future Value:", futureValue)
	// fmt.Println("Future Real Value:", futureRealValue)
	// Printf allows you to format and print a string
	//fmt.Printf("Future Value: %.2f\nFuture Real Value: %.2f", futureValue, futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

//this function accepts three parameters, stores results of calculation in two vars and returns them
func calculateFutureValues (investmentAmount,expectedReturnRate, years float64) (fv float64, frv float64) {
	fv = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	frv = fv / math.Pow(1 + inflationRate / 100, years)
	return fv, frv
}