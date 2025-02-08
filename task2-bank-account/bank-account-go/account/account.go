package account

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	AccountNumber string
	AccountName string
	Balance float64
}	

// NewBankAccount is BankAccount constuctor
func NewBankAccount(accountNumer, AccountName string, initialBalance float64) *BankAccount {
	if initialBalance < 0 {
		initialBalance = 0
	}
	return &BankAccount{
		AccountNumber: accountNumer,
		AccountName: accountNumer,
		Balance: initialBalance,
	}
}

// Deposit adds money to users account
func (b *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater then zero")
	}
	b.Balance += amount
	fmt.Printf("Successfully deposited %.2f. New balance: %.2f\n", amount, b.Balance)
	return nil
}

// Withdraw deducts money from users account
func (b *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be greater than zero")
	}
	if amount > b.Balance {
		return errors.New("insufficient funds")
	}
	b.Balance -= amount
	fmt.Printf("Successfully withdrew %.2f. Remaining balance: %.2f\n", amount, b.Balance)
	return nil
}

// Transfar transfars money to other account
func (b *BankAccount) Transfer(amount float64, receiver *BankAccount) error {
	if amount <= 0 {
		return errors.New("transfer amount must be greater than zero")
	}
	if amount > b.Balance {
		return errors.New("insufficient funds for transfer")
	}
	b.Balance -= amount
	receiver.Balance += amount
	fmt.Printf("Successfully withdrew %.2f. Remaining balance: %.2f\n", amount, b.Balance)
	return nil
}