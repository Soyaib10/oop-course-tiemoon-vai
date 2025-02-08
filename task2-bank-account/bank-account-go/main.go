package main

import (
	"bank-account-go/account"
	"fmt"
)

func main() {
	acc1 := account.NewBankAccount("12345", "Soyaib Rahman", 1000)
	acc2 := account.NewBankAccount("67890", "Rahim Uddin", 500)

	fmt.Println("Initial Balances:")
	fmt.Printf("%s: %.2f\n", acc1.AccountName, acc1.Balance)
	fmt.Printf("%s: %.2f\n", acc2.AccountName, acc2.Balance)

	acc1.Deposit(500)
	acc1.Withdraw(200)
	acc1.Transfer(300, acc2)

	fmt.Println("\nFinal Balances:")
	fmt.Printf("%s: %.2f\n", acc1.AccountName, acc1.Balance)
	fmt.Printf("%s: %.2f\n", acc2.AccountName, acc2.Balance)
}
