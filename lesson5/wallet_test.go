package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want EtherEum) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	assertError := func(t testing.TB, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("wanted an error but didn't get one")
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(EtherEum(10))
		assertBalance(t, wallet, EtherEum(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: EtherEum(20)}
		wallet.Withdraw(EtherEum(10))
		assertBalance(t, wallet, EtherEum(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := EtherEum(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(EtherEum(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsuffientFunds)
	})
}
