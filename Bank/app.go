package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Printf("Your balance is: $%0.2f\n", getBalance(accountBalance))
	} else if choice == 2 {
		fmt.Printf("Your updated account balance is: $%0.2f\n", depositMoney(accountBalance))
	} else if choice == 3 {
		fmt.Printf("Your updated account balance is: $%0.2f\n", withdrawMoney(accountBalance))
	} else if choice == 4 {
		fmt.Println("Goodbye!")
	} else {
		fmt.Println("Invalid choice")
	}
}

func getBalance(accountBalance float64) float64 {
	return accountBalance
}

func depositMoney(accountBalance float64) float64 {
	fmt.Print("Your deposit: ")
	var depositAmount float64
	fmt.Scan(&depositAmount)
	accountBalance += depositAmount
	return accountBalance
}

func withdrawMoney(accountBalance float64) float64 {
	fmt.Print("Your withdrawal: ")
	var withdrawalAmount float64
	fmt.Scan(&withdrawalAmount)
	accountBalance -= withdrawalAmount
	return accountBalance
}
