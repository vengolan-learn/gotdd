package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	wallet := Wallet{}
	(wallet).Deposit(Bitcoin(10))
	got := (wallet).Balance()
	want := Bitcoin(11)

	fmt.Printf("address of balancein test is %v \n", &wallet.balance)
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}

}
