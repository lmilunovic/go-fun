package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin float64

var ErrInsufficientFunds = errors.New("Insufficent funds.")

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(b Bitcoin) error {
	if b > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= b
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%g, BTC", b)
}
