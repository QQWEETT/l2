package main

import "chain/internal"

func main() {
	device := &internal.Device{Name: "Device-1"}
	updateService := &internal.UpdateDataService{Name: "Update-1"}
	dataService := &internal.DataService{}
	device.SetNext(updateService)      // Устройство передает данные сервису обновления
	updateService.SetNext(dataService) // Сервис обновления передает данные сервису сохранения
	data := &internal.Data{}
	device.Execute(data)
}
