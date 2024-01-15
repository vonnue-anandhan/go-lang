package main

import "fmt"

func main() {
	var accountBalance = 1000.0

	for {
		fmt.Println("Welcome to Go bank!")
		fmt.Println("What do you want to do?")

		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)

		case 2:
			fmt.Println("You deposit: ")

			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount // Add money

			fmt.Println("Balance updated! New amount: ", accountBalance)

		case 3:
			fmt.Println("Withdrawal amount: ")

			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. Insufficient balance.")
				continue
			}

			accountBalance -= withdrawalAmount // Deduct money

			fmt.Println("Balance updated! New amount: ", accountBalance)

		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")

			return
		}
	}

}
