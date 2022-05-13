package main

import (
	"builder/internal"
)

func main() {
	asusCollector := internal.GetСollector("asus")
	hpCollector := internal.GetСollector("hp")
	factory := internal.NewFactory(asusCollector) //Создаем завод, который будет производить пк
	asusComputer := factory.CreateComputer()
	asusComputer.Print()
	factory.SetCollector(hpCollector) //Меняем прозводство завода
	hpComputer := factory.CreateComputer()
	hpComputer.Print()
}
