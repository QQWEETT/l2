package main

import (
	"facade/internal"
	"fmt"
)

var (
	bank = internal.Bank{
		Name:  "Тинькофф",
		Cards: []internal.Card{},
	}
	card1 = internal.Card{
		Name:    "CRD-1",
		Balance: 200,
		Bank:    &bank,
	}
	card2 = internal.Card{
		Name:    "CRD-2",
		Balance: 5,
		Bank:    &bank,
	}
	user = internal.User{
		Name: "Покупатель-1",
		Card: &card1,
	}
	user2 = internal.User{
		Name: "Покупатель-2",
		Card: &card2,
	}
	prod = internal.Product{
		Name:  "Сыр",
		Price: 150,
	}
	shop = internal.Shop{
		Name: "Shop",
		Products: []internal.Product{
			prod,
		},
	}
)

func main() {
	fmt.Println("[Банк] Выпуск карт")
	bank.Cards = append(bank.Cards, card1, card2)
	fmt.Printf("[%s]", user.Name)
	err := shop.Sell(user, prod.Name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("[%s]", user2.Name)
	err = shop.Sell(user2, prod.Name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
