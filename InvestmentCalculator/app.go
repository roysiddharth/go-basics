package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5

	var investmentAmount float64
	var years float64
	expectedReturnRate := 7.0

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formatFV := fmt.Sprintf("Future value: %.2f", futureValue)
	formatRealFV := fmt.Sprintf("\nFuture value (adjusted for inflation): %.2f\n", futureRealValue)
	// outputs information
	// fmt.Println("Future Value: ", futureValue)
	// fmt.Println("Future (adjusted for inflation) Value: ", futureRealValue)
	// fmt.Printf("Future value: %.2f", futureValue)
	// fmt.Printf("\nFuture value (adjusted for inflation): %.2f", futureRealValue)
	fmt.Print(formatFV, formatRealFV)
}
