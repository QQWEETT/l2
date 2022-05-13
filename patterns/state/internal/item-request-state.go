package internal

import "fmt"

type itemRequestedState struct {
	vendingMachine *VendingMachine
}

func (i *itemRequestedState) RequestItem() error {
	return fmt.Errorf("Item already requested")
}

func (i *itemRequestedState) AddItem(count int) error {
	return fmt.Errorf("Item Dispense in progress")
}

func (i *itemRequestedState) InsertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		fmt.Errorf("Inserted money is less. Please insert %d", i.vendingMachine.itemPrice)
	}
	fmt.Println("Money entered is ok")
	i.vendingMachine.SetState(i.vendingMachine.hasMoney)
	return nil
}

func (i *itemRequestedState) DispenseItem() error {
	return fmt.Errorf("Please insert money first")
}
