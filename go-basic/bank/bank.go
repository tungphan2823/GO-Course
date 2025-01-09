package main

import "fmt"

func main() {

	var balance = 10000
	fmt.Println("Welcome to TP bank.")
	fmt.Println("What do you want ?")

	fmt.Println("1. Check balance.")
	fmt.Println("2. Deposit money.")
	fmt.Println("3. Withdraw money.")
	fmt.Println("4. Exit.")

	var choice int
	fmt.Println("Your choice: ")
	fmt.Scan(&choice)
	fmt.Println("Your choice: ", choice)

	switch choice {

	case 1:
		fmt.Println("Your balance: ")
		fmt.Println(balance)
	case 2:
		var deposit int
		fmt.Println("Enter amount to deposit: ")
		fmt.Scan(&deposit)
		balance += deposit
	case 3:
		var withdraw int
		fmt.Println("Enter amount to withdraw: ")
		fmt.Scan(&withdraw)
		if withdraw > balance {
			fmt.Println("Insufficient balance.")
		} else {
			balance -= withdraw
		}
	case 4:
		fmt.Println("Thank you for using TP bank.")
		return
	}

}
