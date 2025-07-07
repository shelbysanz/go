package main

import (
	"errors"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line

func updateBalance(customer *customer, transaction transaction) error {
	if transaction.transactionType == "deposit" {
		customer.balance += transaction.amount
		return nil
	} else if transaction.transactionType == "withdrawal" {
		if customer.balance < transaction.amount {
			return errors.New("insufficient funds")
		}
		customer.balance -= transaction.amount
		return nil
	} else {
		return errors.New("unknown transaction type")
	}
}
