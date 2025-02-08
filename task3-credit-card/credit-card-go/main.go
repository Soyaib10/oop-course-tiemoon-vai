package main

import (
	"credit-card-go/creditcard"
	"fmt"
)

func main() {
	card := creditcard.NewCreditCard("Soyaib", "1234")
	
	// per transection
	err := card.WithdrawCash(20000)
	if err != nil {
		fmt.Println(err)
	}

	// paybill transection
	err = card.PayBill(200000)
	err = card.PayBill(300000)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(card.GetCurrentSpending())
}