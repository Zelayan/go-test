package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

    wallet := Wallet{}

    wallet.Deposit(Bitcoin(10))

    got := wallet.Balance()

    want := Bitcoin(110)

    if got != want {
        t.Errorf("got %s want %d", got, want)
    }
}