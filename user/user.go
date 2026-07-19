package user

import (
	"fmt"
	"sync"
)

type User struct {
	ID     string
	Name   string
	Balance float64
	mu      sync.Mutex
}



func (u *User) Deposit(amount float64) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.Balance += amount
	fmt.Println("Депозит выполнен.")
}

func (u *User) Withdraw(amount float64) bool {
	u.mu.Lock()
	defer u.mu.Unlock()
	
	if u.Balance >= amount {
		u.Balance -= amount
		fmt.Println("Снятие успешно выполнено.")
		return true
	}
	fmt.Println("Недостаточно средств для снятия.")
	return false
}