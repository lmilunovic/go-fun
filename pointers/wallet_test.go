package pointers

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet *Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}
	}

	assertError := func(t testing.TB, err error, want error) {
		t.Helper()
		if err == nil {
			t.Fatal("wanted an error but did not get one")
		}

		if err != want {
			t.Errorf("got %q but want %q", err, want)
		}
	}

	assertNoError := func(t testing.TB, err error) {
		if err != nil {
			t.Fatal("Function yields an error")
		}
	}

	t.Run("test bitcoin wallet deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(10.0)

		assertBalance(t, &wallet, want)
	})

	t.Run("test bicoin wallet withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		err := wallet.Withdraw(5)

		want := Bitcoin(5.0)
		assertNoError(t, err)
		assertBalance(t, &wallet, want)
	})

	t.Run("insufficient funds test", func(t *testing.T) {

		wallet := Wallet{}
		startingBalance := Bitcoin(10)
		wallet.Deposit(startingBalance)
		err := wallet.Withdraw(Bitcoin(20))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, &wallet, startingBalance)
	})

}
