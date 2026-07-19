package main

import (
	"fmt"
	"money-transfer/payment"
	"money-transfer/user"
	"sync"
)

func main() {
	// Создаём платёжную систему.
	ps := &payment.PaymentSystem{}

	// Создаём пользователей.
	fmt.Println("Создаю UserID: 1 с балансом 1000")
	fmt.Println("Создаю UserID: 2 с балансом 500")

	user1 := &user.User{
		ID:      "1",
		Name:    "User 1",
		Balance: 1000,
	}

	user2 := &user.User{
		ID:      "2",
		Name:    "User 2",
		Balance: 500,
	}

	// Добавляем пользователей в систему.
	ps.AddUser(user1)
	ps.AddUser(user2)

	// Создаём транзакции.
	fmt.Println("Перевожу с UserID: 1 на UserID: 2 сумму 200")
	fmt.Println("Перевожу с UserID: 2 на UserID: 1 сумму 50")

	t1 := payment.Transaction{
		FromUserID: "1",
		ToUserID:   "2",
		Amount:     200,
	}

	t2 := payment.Transaction{
		FromUserID: "2",
		ToUserID:   "1",
		Amount:     50,
	}

	// Добавляем транзакции в очередь.
	ps.AddTransaction(t1)
	ps.AddTransaction(t2)

	// Создаём буферизированный канал.
	ch := make(
		chan payment.Transaction,
		len(ps.TransactionQueue),
	)

	// Создаём WaitGroup.
	var wg sync.WaitGroup

	// Запускаем три воркера.
	workersCount := 3

	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go ps.Worker(ch, &wg)
	}

	// Отправляем все транзакции в канал.
	for _, transaction := range ps.TransactionQueue {
		ch <- transaction
	}

	// Сообщаем воркерам, что новых транзакций не будет.
	close(ch)

	// Ждём завершения всех воркеров.
	wg.Wait()

	// Выводим итоговые балансы.
	fmt.Println("Итого")
	fmt.Printf(
		"У первого пользователя должно получиться 850, получилось %.2f\n",
		user1.Balance,
	)
	fmt.Printf(
		"У второго пользователя должно получиться 650, получилось %.2f\n",
		user2.Balance,
	)
}