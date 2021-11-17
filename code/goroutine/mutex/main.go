package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

type Account struct {
	Balance int
}

func DepositAndWithdraw(account *Account) {
	mutex.Lock()         // Mutex Key 획득
	defer mutex.Unlock() // defer를 사용한 Unlock()
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}

func main() {
	var wg sync.WaitGroup

	account := &Account{0}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			// for {
			DepositAndWithdraw(account)
			// }
		} ()
		wg.Done()
	}
	wg.Wait()
}