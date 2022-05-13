package internal

import (
	"errors"
	"fmt"
	"time"
)

//Создаем тип банка
type Bank struct {
	Name  string
	Cards []Card
}

//Проверяем, положительный баланс или нет
func (bank Bank) CheckBalance(cardNumber string) error {
	fmt.Printf("[Банк] Получение остатка по карте %s\n", cardNumber)
	time.Sleep(time.Millisecond * 500)
	//Ищем карту, по которой происходит запрос
	for _, card := range bank.Cards {
		if card.Name != cardNumber {
			continue
		}
		if card.Balance <= 0 {
			return errors.New("[Банк] Недостаточно средств!")
		}
	}
	fmt.Println("[Банк] Остаток положительный!")
	return nil
}
