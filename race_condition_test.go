package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {

	var mutex sync.Mutex
	var x int32 = 0
	// theres 1000 goroutine
	for i := 0; i < 1000; i++ {
		go func() {
			// each will try to increment x 100 times
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x += 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)

	// in total there sould be 1000 * 100 = 100000 increment
	fmt.Println(x)

}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance uint64
}

func (account *BankAccount) AddBalance(amount uint64) {
	account.RWMutex.Lock()
	defer account.RWMutex.Unlock()
	account.Balance += amount
}

func (account *BankAccount) GetBalance() uint64 {
	account.RWMutex.RLock()
	defer account.RWMutex.RUnlock()
	balance := account.Balance // prevent race condition reading
	return balance
}

func TestRaceConditionWithRWMutex(t *testing.T) {
	someUserBankAccount := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				someUserBankAccount.AddBalance(1000)
				fmt.Println("Now your balance is:", someUserBankAccount.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Final balance is:", someUserBankAccount.GetBalance()) // should be 10000000
}
