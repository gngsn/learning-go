package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	/*

		1) got := wallet.Balance()
		2) fmt.Printf("address of balance in test is %v \n", &wallet.balance)

		1) address of balance in Deposit is 0xc0000b01a8 
		2) address of balance in test    is 0xc0000b01a0

		!! Method는 호출될 때 Copy된다.
		처음 생성될 때와 호출될 때의 wallet의 balance는 다른 메모리를 가진 다른 데이터 
	*/

    t.Run("Deposit", func(t *testing.T) {
        wallet := Wallet{}
        wallet.Deposit(Bitcoin(10))

        assertBalance(t, wallet, Bitcoin(10))
    })

    t.Run("Withdraw with funds", func(t *testing.T) {
        wallet := Wallet{Bitcoin(20)}
        err := wallet.Withdraw(Bitcoin(10))

        assertNoError(t, err)
        assertBalance(t, wallet, Bitcoin(10))
    })

    t.Run("Withdraw insufficient funds", func(t *testing.T) {
        wallet := Wallet{Bitcoin(20)}
        err := wallet.Withdraw(Bitcoin(100))

        assertError(t, err, ErrInsufficientFunds)
        assertBalance(t, wallet, Bitcoin(20))
    })
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
    t.Helper()
    got := wallet.Balance()

    if got != want {
        t.Errorf("got %s want %s", got, want)
    }
}

/* 
	'왜 안되지..?'의 경우
	원하지 않은 Error Assert
*/
func assertNoError(t testing.TB, got error) {
    t.Helper()
    if got != nil {
        t.Fatal("got an error but didn't want one")
    }
}

/* 
	'왜 되지..?'의 경우
	Error를 원하는데 없을 때 Catch
*/
func assertError(t testing.TB, got error, want error) {
    t.Helper()
    if got == nil {
        t.Fatal("didn't get an error but wanted one")
    }

    if got != want {
        t.Errorf("got %s, want %s", got, want)
    }
}