package main

import (
	"bank/fileops"
	"fmt"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFileName = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFileName)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------------")

		// return
		// panic("Can't continue, sorry")
	}

	fmt.Println("Welcome to Go bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	for {
		presentOptions()

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
			fileops.WriteFloatToFile(accountBalanceFileName, accountBalance)

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
			fileops.WriteFloatToFile(accountBalanceFileName, accountBalance)

			fmt.Println("Balance updated! New amount: ", accountBalance)

		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")

			return
		}
	}

}
