package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	expenses, err1 := getUserInput("Expenses: ")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	taxRate, err2 := getUserInput("Tax Rate: ")
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)

	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
	writeDataToFile(ebt, profit, ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func writeDataToFile(ebt, profit, ratio float64) {
	ebtFixed := fmt.Sprint(ebt)
	profitFixed := fmt.Sprint(profit)
	ratioFixed := fmt.Sprint(ratio)

	data := fmt.Sprintf("ebt is %s, profit is %s, ratio is %s.\n", ebtFixed, profitFixed, ratioFixed)
	os.WriteFile("data.txt", []byte(data), 0644)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return userInput, errors.New("invalid input. please enter a positive number.")
	}
	return userInput, nil
}
