package internal

import "fmt"

type NoItemState struct {
	vendingMachine *VendingMachine
}

//Запрос товара
func (i *NoItemState) RequestItem() error {
	return fmt.Errorf("Item out of stock")
}

//Добавление товара
func (i *NoItemState) AddItem(count int) error {
	i.vendingMachine.IncrementItemCount(count)
	i.vendingMachine.SetState(i.vendingMachine.hasItem)
	return nil
}

//Внесение денег
func (i *NoItemState) InsertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

//Выдача товара
func (i *NoItemState) DispenseItem() error {
	return fmt.Errorf("Item out of stock")
}
