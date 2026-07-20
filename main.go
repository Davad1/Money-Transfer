package main

import (
	"fmt"
	"money-transfer/payment"
	"money-transfer/user"
	"sync"
)

func main() {

	ps := &payment.PaymentSystem{}


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

	
	ps.AddUser(user1)
	ps.AddUser(user2)


	// Подсказка
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

	ps.AddTransaction(t1)
	ps.AddTransaction(t2)

	
	ch := make(
		chan payment.Transaction,
		len(ps.TransactionQueue),
	)


	var wg sync.WaitGroup


	workersCount := 3

	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go ps.Worker(ch, &wg)
	}

	
	for _, transaction := range ps.TransactionQueue {
		ch <- transaction
	}

	close(ch)


	wg.Wait()

	// Подсказка
	fmt.Println("Итого")
	fmt.Printf("User 1 Баланс: %v", user1.Balance)
	fmt.Printf("User 2 Баланс: %v", user2.Balance)
}