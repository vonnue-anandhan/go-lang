package main

import (
	"fmt"
	"math"
)

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	const inflationRate = 2.5

	futureValue := investmentAmount * math.Pow((1+(expectedReturnRate)/100), years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	return futureValue, futureRealValue
}

func main() {

	var investmentAmount, years float64
	expectedReturnRate := 5.5

	outputText("Investment amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Years: ")
	fmt.Scan(&years)

	outputText("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future value (adjusted for inflation): %.1f\n", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)

	// fmt.Print(`Hello
	// World
	// `)

}
