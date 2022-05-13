package internal

import "fmt"

type hasItemState struct {
	vendingMachine *VendingMachine
}

func (i *hasItemState) RequestItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.SetState(i.vendingMachine.noItem)
		return fmt.Errorf("No item present")
	}
	fmt.Printf("Item requestd\n")
	i.vendingMachine.SetState(i.vendingMachine.itemRequested)
	return nil
}

func (i *hasItemState) AddItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.vendingMachine.IncrementItemCount(count)
	return nil
}

func (i *hasItemState) InsertMoney(money int) error {
	return fmt.Errorf("Please select item first")
}
func (i *hasItemState) DispenseItem() error {
	return fmt.Errorf("Please select item first")
}
