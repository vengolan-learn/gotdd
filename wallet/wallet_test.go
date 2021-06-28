package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	wallet := Wallet{}
	(wallet).Deposit(10)
	got := (wallet).Balance()
	want := 10

	fmt.Printf("address of balancein test is %v \n", &wallet.balance)
	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}

}
