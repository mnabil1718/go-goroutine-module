package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (userBalance *UserBalance) Lock() {
	userBalance.Mutex.Lock()
}

func (userBalance *UserBalance) Unlock() {
	userBalance.Mutex.Unlock()
}

func (userBalance *UserBalance) ChangeBalance(amount int) {
	userBalance.Balance += amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Locking user1 " + user1.Name + "...")
	user1.ChangeBalance(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Locking user2 " + user2.Name + "...")
	user2.ChangeBalance(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadLock(t *testing.T) {

	user1 := UserBalance{Name: "Budi", Balance: 1000}
	user2 := UserBalance{Name: "Agus", Balance: 1000}

	go Transfer(&user1, &user2, 1000)
	go Transfer(&user2, &user1, 1000)

	time.Sleep(10 * time.Second)

	fmt.Println("User 1 balance:", user1.Balance)
	fmt.Println("User 2 balance:", user2.Balance)

}
