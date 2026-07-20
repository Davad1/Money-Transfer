package payment


import (
	"fmt"
	"money-transfer/user"
)

type PaymentSystem struct {
	users map[string]*user.User
	TransactionQueue []Transaction
}


func (ps *PaymentSystem) AddUser(u *user.User) {
	if ps.users == nil {
		ps.users = make(map[string]*user.User)
	}
	ps.users[u.ID] = u
}

func (ps *PaymentSystem) AddTransaction(t Transaction) {
	if ps.TransactionQueue == nil {
		ps.TransactionQueue = make([]Transaction, 0, 10)
	}
	ps.TransactionQueue = append(ps.TransactionQueue, t)
}

func (ps *PaymentSystem) ProcessingTransactions() error {
	for _, t := range ps.TransactionQueue {

		// Ищем пользователя, который отправляет деньги
		fromUser, ok := ps.users[t.FromUserID]
		if !ok {
			return fmt.Errorf("пользователь %s не найден", t.FromUserID)
		}

		// Ищем пользователя, который получает деньги
		toUser, ok := ps.users[t.ToUserID]
		if !ok {
			return fmt.Errorf("пользователь %s не найден", t.ToUserID)
		}

		// Снимаем деньги
		if err := fromUser.Withdraw(t.Amount); err != nil {
			return fmt.Errorf("ошибка снятия у пользователя %s: %w", t.FromUserID, err)
		}

		// Начисляем деньги
		toUser.Deposit(t.Amount)
	}

	return nil

}