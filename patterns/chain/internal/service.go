package internal

type Service interface {
	Execute(*Data)
	SetNext(Service)
}

type Data struct {
	GetSource    bool //Проверяет, выполнился ли прием данных
	UpdateSource bool //Ставит отметку тот сервис, который обработал данные
}
