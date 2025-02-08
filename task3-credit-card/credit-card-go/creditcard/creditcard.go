package creditcard

import (
	"errors"
	"fmt"
	"time"
)

// CreditCard struct with account name & limits
type CreditCard struct {
	accountName      string
	cardNumber       string
	creditLimit      float64
	currentSpending  float64
	dailyWithdrawal  float64
	lastWithdrawDate time.Time
}

// Constructor to create a new Credit Card
func NewCreditCard(accountName, cardNumber string) *CreditCard {
	return &CreditCard{
		accountName:      accountName,
		cardNumber:       cardNumber,
		creditLimit:      500000,
		currentSpending:  0,
		dailyWithdrawal:  0,
		lastWithdrawDate: time.Now(),
	}
}

// ResetDailyLimitIfNeeded resets daily limit
func (c *CreditCard) ResetDailyLimitIfNeeded() {
	today := time.Now().Format("2006-01-02")
	lastDay := c.lastWithdrawDate.Format("2006-01-02")

	if today != lastDay {
		c.dailyWithdrawal = 0
		c.lastWithdrawDate = time.Now()
	}
}

// WithdrawCash withdraw money for per, daily, total transection
func (c *CreditCard) WithdrawCash(amount float64) error {
	c.ResetDailyLimitIfNeeded()
	if amount > 20000 {
		return errors.New("per transection limit exceeded")
	}
	if c.dailyWithdrawal + amount > 100000 {
		return errors.New("daily withdrawal limit exceeded")
	}
	if c.currentSpending + amount > c.creditLimit {
		return errors.New("credit limit exceeded")
	}

	c.currentSpending += amount
	c.dailyWithdrawal += amount
	fmt.Printf("%s withdraw %.2f\n", c.accountName, amount)
	return nil
}

// PayBill checks total credit limit
func (c *CreditCard) PayBill(amount float64) error {
	if c.currentSpending + amount > c.creditLimit {
		return errors.New("credit limit exceeded")
	}
	c.currentSpending += amount
	fmt.Printf("%s paid a bill of: %.2f\n", c.accountName, amount)
	return nil
}

func (c *CreditCard) GetCurrentSpending() float64 {
	return c.currentSpending
}