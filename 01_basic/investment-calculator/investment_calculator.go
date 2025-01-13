package main

import (
	"fmt"
	"math"
)

const inflations = 2.5

func main() {

	var investmentAmount float64
	fmt.Print("Enter your initial investment amount: $")
	fmt.Scan(&investmentAmount)

	var expectedReturnRate = 5.5
	var years float64 = 10
	futureValue, futureRealMoney := calculateFutureValue(investmentAmount, years, expectedReturnRate)
	fmt.Printf("Future Value: $%.2f\n", futureValue)
	fmt.Printf("Future Real Money: $%.2f\n", futureRealMoney)

}
func calculateFutureValue(investmentAmount float64, years float64, expectedReturnRate float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflations/100, years)
	return fv, rfv
}
