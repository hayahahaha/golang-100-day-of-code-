package poiters

import (
	"math/rand"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Test wallet balance", func(t *testing.T) {
		amount := rand.Intn(100)
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(amount))
		want := Bitcoin(amount)

		assertBalance(t, wallet, want)
	})
	t.Run("Test withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(40)}
		err := wallet.Withdraw(10)
		want := Bitcoin(30)

		assertNoError(t, err)
		assertBalance(t, wallet, want)
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		got := wallet.Withdraw(Bitcoin(50))
		want := "cannot withdraw, insufficient funds"

		assertError(t, got, want)
		assertBalance(t, wallet, startingBalance)

	})
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("Got an error but dont want one")
	}
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want string) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if got.Error() != want {
		t.Errorf("got %q want %q", got, want)
	}
}
