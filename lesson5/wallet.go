package main

import (
	"errors"
	"fmt"
)

type Stringer interface {
	String() string
}

// Crypto

type EtherEum int

// ^ spelling creds to LukeOfLoxley

func (e EtherEum) String() string {
	return fmt.Sprintf("%d ETH", e)
}

// Wallet

type Wallet struct {
	balance EtherEum
}

func (w *Wallet) Deposit(amount EtherEum) {
	w.balance += amount
}

var ErrInsuffientFunds = errors.New("Cannot withdraw, insuffient funds")

func (w *Wallet) Withdraw(amount EtherEum) error {
	if amount > w.balance {
		return ErrInsuffientFunds
	}

	w.balance -= amount
	return nil
}

func (w Wallet) Balance() EtherEum {
	return w.balance
}
