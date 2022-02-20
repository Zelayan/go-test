package pointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWallet(t *testing.T) {


	check := assert.New(t)
    assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
        got := wallet.Balance()

        if got != want {
            t.Errorf("got %s want %s", got, want)
        }
    }

    t.Run("Deposit", func(t *testing.T) {
        wallet := Wallet{}
        wallet.Deposit(Bitcoin(10))
        assertBalance(t, wallet, Bitcoin(10))
    })

    t.Run("Withdraw", func(t *testing.T) {
        wallet := Wallet{balance: Bitcoin(20)}
        err := wallet.Withdraw(Bitcoin(10))
        assertBalance(t, wallet, Bitcoin(10))
		check.Nil(err)
    })

	t.Run("Withdraw more than", func(t *testing.T) {
        wallet := Wallet{balance: Bitcoin(20)}
        err := wallet.Withdraw(Bitcoin(30))
        assertBalance(t, wallet, Bitcoin(10))
		check.Nil(err)
    })

}