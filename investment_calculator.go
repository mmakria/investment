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

	fmt.Print("Enter an investment Amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter years: ")
	fmt.Scan(&years)

	futureValue := (investmentAmount) * math.Pow((1+expectedReturnRate/100), years)

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Printf("Future value is %0.2f \nFuture value including inflation is %0.2f ", futureValue, futureRealValue)

}
