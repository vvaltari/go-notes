package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount, expectedReturnRate, years float64

	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calcFutureValue(investmentAmount, expectedReturnRate, years)

	fmt.Printf("The future value would be: %.2f\n", futureValue)
	fmt.Printf("The real future value would be: %.2f", futureRealValue)
}

func calcFutureValue(investmentAmount, expectedReturnRate, years float64) (futureValue float64, futureRealValue float64) {
	futureValue = investmentAmount	* math.Pow(1 + expectedReturnRate / 100, years)
	futureRealValue = futureValue / math.Pow(1 + inflationRate / 100, years)

	return futureValue, futureRealValue
}