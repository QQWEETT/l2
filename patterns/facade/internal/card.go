package internal

import (
	"fmt"
	"time"
)

//Создаем тип карт
type Card struct {
	Name    string
	Balance float64
	Bank    *Bank
}

func (card Card) CheckBalance() error {
	fmt.Println("[Карта] Запрос в банк для проверки остатка")
	time.Sleep(500 * time.Millisecond)
	return card.Bank.CheckBalance(card.Name)

}
