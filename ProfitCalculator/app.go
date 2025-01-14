package main

import "fmt"

const taxRate float64 = 0.3

func main() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")

	ebt, profit, ratio := calculateProfit(revenue, expenses)

	fmt.Printf("\nEBT: %v\n", ebt)
	fmt.Printf("Profit: %v\n", profit)
	fmt.Printf("Ratio: %0.2f\n", ratio)
}

func getUserInput(inputText string) (userInput float64) {
	fmt.Print(inputText)
	fmt.Scan(&userInput)
	return userInput
}

func calculateProfit(revenue float64, expenses float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate)
	ratio := ebt / profit

	return ebt, profit, ratio
}
