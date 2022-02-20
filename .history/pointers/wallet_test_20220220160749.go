package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

    wallet := Wallet{}

    //fmt.Println("address of balance in test is", &wallet.balance)

    got := wallet.Balance()
    want := 10

    if got != want {
        t.Errorf("got %d want %d", got, want)
    }
}