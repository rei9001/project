package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	Balance int
}

func main() {

	var wg sync.WaitGroup

	account := &Account{10}

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; i < 10; j++ {
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func DepositAndWithdraw(account *Account) {
	var mutex sync.Mutex

	mutex.Lock()

	defer mutex.Unlock()

	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance shoud not be negative value: %d", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}
