package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()
		if got != want {
			t.Errorf("want %s, got %s", want, got)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		(wallet).Deposit(Bitcoin(10))
		want := Bitcoin(10)

		fmt.Printf("address of balancein test is %v \n", &wallet.balance)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraow", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(8))
		want := Bitcoin(12)
		assertBalance(t, wallet, want)
	})

}
