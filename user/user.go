package user

import (
	"fmt"
	
)

type User struct {
	ID     string
	Name   string
	Balance float64
	
}


func (u *User) Deposit(amount float64) {
	u.Balance += amount
	fmt.Println("Депозит выполнен.")
}

func (u *User) Withdraw(amount float64) bool {

	if u.Balance >= amount {
		u.Balance -= amount
		fmt.Println("Снятие успешно выполнено.")
		return true
	}
	fmt.Println("Недостаточно средств для снятия.")
	return false
}