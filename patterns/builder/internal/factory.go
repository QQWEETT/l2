package internal

type Factory struct {
	Collector Collector //Реализуем поле интерфейса сборки
}

//Передаем интерфейс, как аргумент, чтобы мы могли задавать любую комплектацию пк,
//которую мы создаем на заводе
func NewFactory(collector Collector) *Factory {
	return &Factory{Collector: collector}
}

//Создаем функцию завода, чтобы мы могли менять в любое время комплектацию
func (factory *Factory) SetCollector(collector Collector) {
	factory.Collector = collector
}

//Создаем функцию, которая будет создавать компьютеры
func (factory *Factory) CreateComputer() Computer {
	factory.Collector.SetCore()
	factory.Collector.SetMemory()
	factory.Collector.SetBrand()
	factory.Collector.SetGraphicCard()
	factory.Collector.SetMonitor()
	return factory.Collector.GetComputer()
}
