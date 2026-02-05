package main

import (
	"fmt"
)

func main() {
	// Hardcoded PIN and starting balance
	const correctPIN = 1234
	balance := 1000.0

	var pin int
	fmt.Println("Welcome to Mini ATM!")

	// Ask for PIN
	fmt.Print("Enter your 4-digit PIN: ")
	fmt.Scan(&pin)
	if pin != correctPIN {
		fmt.Println("Incorrect PIN! Exiting...")
		return
	}

	// ATM menu loop
	for {
		fmt.Println("\nATM Menu:")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your balance is: $%.2f\n", balance)
		case 2:
			var deposit float64
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&deposit)
			if deposit > 0 {
				balance += deposit
				fmt.Printf("Deposited $%.2f. New balance: $%.2f\n", deposit, balance)
			} else {
				fmt.Println("Invalid amount!")
			}
		case 3:
			var withdraw float64
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&withdraw)
			if withdraw > 0 && withdraw <= balance {
				balance -= withdraw
				fmt.Printf("Withdrew $%.2f. New balance: $%.2f\n", withdraw, balance)
			} else {
				fmt.Println("Insufficient balance or invalid amount!")
			}
		case 4:
			fmt.Println("Thank you for using Mini ATM. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}
