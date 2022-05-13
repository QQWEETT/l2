package internal

import "fmt"

type Device struct {
	Name string
	Next Service
}

func (device *Device) Execute(data *Data) {
	//Проверяем состояние данных
	if data.GetSource {
		fmt.Printf("Data from device [%s] already get.\n", device.Name)
		device.Next.Execute(data)
		return
	}
	fmt.Printf("Get data from device [%s].\n", device.Name)
	data.GetSource = true
	device.Next.Execute(data)
}

// Задаем следующее звено
func (device *Device) SetNext(svc Service) {
	device.Next = svc
}
