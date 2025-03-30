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

	futureValue := investmentAmount	* math.Pow(1 + expectedReturnRate / 100, years)
	futureValue = futureValue / math.Pow(1 + inflationRate / 100, years)

	fmt.Println("The future value would be:", futureValue)
}