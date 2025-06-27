package main

import (
	"fmt"
	"math"
)

func main() {
	profitCalc()
}

func profitCalc() {
	//Variable declaration
	var revenue, expenses, taxRate float64

	//Actual function
	fmt.Print("What is your expected revenue? ")
	fmt.Scan(&revenue)

	fmt.Print("What are your expenses? ")
	fmt.Scan(&expenses)

	fmt.Print("What is the tax rate? ")
	fmt.Scan(&taxRate)

	//Calculation
	EBT := revenue-expenses
	profit := (revenue-expenses)*((100-taxRate)/100)
	
	fmt.Println("Revenue entered: ", revenue)
	fmt.Println("Expenses entered: ", expenses)
	fmt.Println("Tax rate entered: ", taxRate)
	fmt.Println("Earnings before tax (EBT): ", math.Round(EBT))
	fmt.Println("Total profit: ", math.Round(profit))
}