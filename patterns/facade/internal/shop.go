package internal

import (
	"errors"
	"fmt"
	"time"
)

//Создаем тип магазина
type Shop struct {
	Name     string
	Products []Product
}

//Создаем фасад
func (shop Shop) Sell(user User, product string) error {
	fmt.Println("[Магазин] Запрос к пользователю для получения остатка по карте")
	time.Sleep(time.Millisecond * 500)
	//Проверяем баланс пользователя
	err := user.Card.CheckBalance()
	if err != nil {
		return err
	}
	fmt.Printf("[Магазин] Проверка - может ли [%s] купить товар \n", user.Name)
	time.Sleep(time.Millisecond * 500)
	for _, prod := range shop.Products {
		if prod.Name != product {
			continue
		}
		if prod.Price > user.GetBalance() {
			return errors.New("[Магазин] Недостаточно средств для покупки товара!")
		}
		fmt.Printf("[Магазин] Товар [%s] куплен \n", prod.Name)

	}
	return nil
}
