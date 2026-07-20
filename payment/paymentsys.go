package payment


import (
	"fmt"
	"money-transfer/user"
	"sync"
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

func (ps *PaymentSystem) ProcessTransaction(t Transaction) error {
	fromUser, ok := ps.users[t.FromUserID]
	if !ok {
		return fmt.Errorf(
			"пользователь %s не найден",
			t.FromUserID,
		)
	}

	toUser, ok := ps.users[t.ToUserID]
	if !ok {
		return fmt.Errorf("Пользователь %s не найден", t.ToUserID,)
	}

	err := fromUser.Withdraw(t.Amount)
	if err != nil {
		return fmt.Errorf("ошибка снятия у пользователя %v: %v", t.FromUserID, err)
	}

	toUser.Deposit(t.Amount)

	return nil
}


func (ps *PaymentSystem) Worker(ch <-chan Transaction, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for t := range ch {
		err := ps.ProcessTransaction(t)
		if err != nil {
			fmt.Println("Ошибка при обработке транзакции:", err)
			continue
		}

		fmt.Printf("Транзакция успешно обработана: %s -> %s, сумма: %.2f\n", t.FromUserID, t.ToUserID, t.Amount)
	}
}