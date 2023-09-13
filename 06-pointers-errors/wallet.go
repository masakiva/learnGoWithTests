package main

import (
	"errors" // https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
} // https://golang.org/pkg/fmt/#Stringer

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount // struct pointers are automatically dereferenced https://golang.org/ref/spec#Method_values
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
