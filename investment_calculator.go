package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	//fmt.Print("Enter an investment Amount: ")
	outputText("Enter an investment Amount: ")
	fmt.Scan(&investmentAmount)
	//fmt.Print("Enter years: ")
	outputText("Enter Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, years, expectedReturnRate)

	//futureValue := (investmentAmount) * math.Pow(1+expectedReturnRate/100, years)
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFv := fmt.Sprintf("Future value is %0.2f \n", futureValue)
	formattedRfv := fmt.Sprintf("Future real value is %0.2f \n", futureRealValue)
	// fmt.Printf("Future value is %0.2f \nFuture value including inflation is %0.2f ", futureValue, futureRealValue)

	fmt.Print(formattedFv, formattedRfv)
	outputText(formattedFv)
	outputText(formattedRfv)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount,
	expectedReturnRate,
	years float64) (fv float64, frv float64) {
	fv = (investmentAmount) * math.Pow(1+expectedReturnRate/100, years)
	frv = fv / math.Pow(1+inflationRate/100, years)
	return fv, frv
	//return
}
