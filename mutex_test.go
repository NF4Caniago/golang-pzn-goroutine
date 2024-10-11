package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x++
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("x = ", x)
}

type BackAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BackAccount) addBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BackAccount) getBalance() int {
	account.RWMutex.RLock()
	ballance := account.Balance
	account.RWMutex.RUnlock()
	return ballance
}

func TestRWMutex(t *testing.T) {
	account := BackAccount{}
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.addBalance(1000)
				fmt.Println(account.getBalance())
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Total Ballance : ", account.getBalance())
}
