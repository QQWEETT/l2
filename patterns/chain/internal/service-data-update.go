package internal

import "fmt"

type UpdateDataService struct {
	Name string
	Next Service
}

func (upd *UpdateDataService) Execute(data *Data) {
	//Проверяем состояние данных
	if data.UpdateSource {
		fmt.Printf("Data in service [%s] is already update.\n", upd.Name)
		upd.Next.Execute(data)
		return
	}
	fmt.Printf("Update data from service [%s].\n", upd.Name)
	data.UpdateSource = true
	upd.Next.Execute(data)
}

// Задаем следующее звено
func (upd *UpdateDataService) SetNext(svc Service) {
	upd.Next = svc
}
