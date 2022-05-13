package internal

import "fmt"

type hasMoneyState struct {
	vendingMachine *VendingMachine
}

func (i *hasMoneyState) RequestItem() error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *hasMoneyState) AddItem(count int) error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *hasMoneyState) InsertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

func (i *hasMoneyState) DispenseItem() error {
	fmt.Println("Dispensing Item")
	i.vendingMachine.itemCount = i.vendingMachine.itemCount - 1
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.SetState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.SetState(i.vendingMachine.hasItem)
	}
	return nil
}
