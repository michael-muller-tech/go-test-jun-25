package main

import (
	"fmt"
	"math"
)

func main() {
	//Core values
	const inflationRate = 2.5
	var investmentAmount float64
	years := 10.0
	expectedReturnRate := 5.5


	//User input
	fmt.Print("Please enter a value for investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Please enter the number of years for your investment: ")
	fmt.Scan(&years)

	fmt.Print("Please enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	
	//Calculations
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100,years)
	fmt.Println(math.Round(futureValue))
	fmt.Println(math.Round(futureRealValue))
}