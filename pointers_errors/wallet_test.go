package main

import (
	"errors"
	"fmt"
	"testing"
)

var ErrorInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int
type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	//fmt.Printf("the address is %v\n",&w.balance)
	w.balance += amount
}
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrorInsufficientFunds
	}
	w.balance -= amount
	return nil
}
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
func TestWallet(t *testing.T) {
	t.Run("Depisot", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertion(t, wallet, Bitcoin(10))
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoerror(t, err)
		assertion(t, wallet, Bitcoin(10))
	})
	t.Run("Withdraw", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(30))
		assertion(t, wallet, Bitcoin(20))
		//fmt.Println("cccccccca")
		assertError(t, err, ErrorInsufficientFunds)
		//fmt.Println("cccccccc")
	})
}
func assertion(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s but want %s", got, want)
	}
}
func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but did not get one")
	}
	if got != want {
		t.Errorf("got %q,but want %q", got, want)
	}
}
func assertNoerror(t *testing.T, err error) {
	if err != nil {
		t.Fatal("get an error but did not want one!")
	}
}
