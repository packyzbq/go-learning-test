package pointer

import (
	"errors"
	"fmt"
)

type Coin int

type Wallet struct {
	ballance Coin
}

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

type WalletInt struct {
	ballance int
}

func (w WalletInt) Deposit(amount int) {
	w.ballance += amount
}

func (w WalletInt) Ballance() int {
	return w.ballance
}

func (w *Wallet) DepositP(amount Coin) {
	w.ballance += amount
}

func (w *Wallet) BallanceP() Coin {
	return w.ballance
}

func (w *Wallet) WithdrawP(amount Coin) error {
	if amount > w.ballance {
		return InsufficientFundsError
	}

	w.ballance -= amount
	return nil
}

func (c Coin) String() string {
	return fmt.Sprintf("%d Coin", c)
}
