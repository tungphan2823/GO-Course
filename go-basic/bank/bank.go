package main

import (
	"fmt"

	"example.com/bank/fileops"
)

const textFile = "balance.txt"

func main() {

	var balance, err = fileops.ReadDataFromFile(textFile)
	if err != nil {
		fmt.Printf("ERROR")
		fmt.Println(err)
		// panic("can't continue")
	}
	fmt.Println("Welcome to TP bank.")

	for {
		presentOptions()

		var choice int
		fmt.Println("Your choice: ")
		fmt.Scan(&choice)
		fmt.Println("Your choice: ", choice)

		switch choice {

		case 1:
			fmt.Println("Your balance: ")
			fmt.Println(balance)
		case 2:
			var deposit float64
			fmt.Println("Enter amount to deposit: ")
			fmt.Scan(&deposit)
			if deposit < 0 {
				fmt.Println("Invalid deposit amount.")
				continue
			}
			balance += deposit
			fileops.WriteDataToFile(float64(balance), textFile)
		case 3:
			var withdraw float64
			fmt.Println("Enter amount to withdraw: ")
			fmt.Scan(&withdraw)
			if withdraw > balance {
				fmt.Println("Insufficient balance.")
				continue
			} else {
				balance -= withdraw
				fileops.WriteDataToFile(float64(balance), textFile)
			}
		default:
			fmt.Println("Thank you for using TP bank.")
			return
		}
	}

}
