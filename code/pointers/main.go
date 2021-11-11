package pointers

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

type Wallet struct {
    balance Bitcoin
}

/*
	func (w Wallet) Deposit(amount int) {
		fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
		w.balance += amount
	}

	func (w Wallet) Balance() int {
		return w.balance
	}

	-> Method는 호출될 때 Copy되기 때문에 pointer를 사용해서 해결해보자 ~
*/


func (w *Wallet) Deposit(amount Bitcoin) {
    w.balance += amount
}


func (w *Wallet) Balance() Bitcoin {
	/*
    	return (*w).balance

		사실 위의 코드가 정확하지만, Go 개발자들은 위의 표기가 번거롭다고 생각해서 dereference과정을 명시적으로 표시하지 않아도 수행 되도록 제작함.
		실제로 위와 아래 코드는 동일하게 수행됨
	*/
    return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

    if amount > w.balance {
        return ErrInsufficientFunds
    }

    w.balance -= amount
    return nil
}

/* fmt package 에서 제공하는 interface

	type Stringer interface {
			String() string
	}
*/
func (b Bitcoin) String() string {
    return fmt.Sprintf("%d BTC", b)
}