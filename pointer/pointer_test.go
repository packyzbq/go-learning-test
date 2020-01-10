package pointer

import "testing"

func TestWallet(t *testing.T) {

	wallet := WalletInt{}

	wallet.Deposit(10)

	got := wallet.Ballance()
	want := 10

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestWalletP(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Coin) {
		got := wallet.BallanceP()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t *testing.T, got error, want string) {
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}

		if got.Error() != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.DepositP(Coin(10))
		assertBalance(t, wallet, Coin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{ballance: Coin(20)}
		wallet.WithdrawP(Coin(10))
		assertBalance(t, wallet, Coin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Coin(20)
		wallet := Wallet{startingBalance}
		err := wallet.WithdrawP(Coin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, "cannot withdraw, insufficient funds")

	})

}
