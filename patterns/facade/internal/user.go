package internal

//Создаем пользователя
type User struct {
	Name string
	Card *Card
}

//Создаем функцию, которая возвращает баланс
func (user User) GetBalance() float64 {
	return user.Card.Balance
}
