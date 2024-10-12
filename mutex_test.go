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

type UserBallance struct {
	sync.Mutex
	Name     string
	Ballance int
}

func (user *UserBallance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBallance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBallance) Change(amount int) {
	user.Ballance = user.Ballance + amount
}

func Transfer(user1 *UserBallance, user2 *UserBallance, amount int) {
	user1.Lock()
	fmt.Println("Lock User 1 : ", user1.Name)
	user1.Change(-amount)
	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock User 2 : ", user2.Name)
	user2.Change(amount)
	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadLock(t *testing.T) {
	user1 := UserBallance{
		Name:     "Afif",
		Ballance: 10000,
	}

	user2 := UserBallance{
		Name:     "Ilham",
		Ballance: 20000,
	}

	go Transfer(&user1, &user2, 1000)
	go Transfer(&user2, &user1, 2000)

	time.Sleep(7 * time.Second)

	fmt.Println("User 1 Name : ", user1.Name, ", Ballance : ", user1.Ballance)
	fmt.Println("User 2 Name : ", user2.Name, ", Ballance : ", user2.Ballance)
}
